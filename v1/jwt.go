package pluginsv1

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// JWTClaims holds the verified identity extracted from a validated JWT.
type JWTClaims struct {
	Subject   string
	Email     string
	Roles     []string
	// SessionID is the value of the "sid" claim, if present.
	// Used for session revocation: after RevokeSession is called with this ID,
	// ValidateToken will return an error for any token carrying the same sid.
	SessionID string
	// ExpiresAt is the token expiry (unix seconds), used to auto-evict
	// the revocation entry once the token can no longer be used anyway.
	ExpiresAt int64
}

// JWTValidator verifies RS256 JWTs against a JWKS endpoint and enforces
// session revocation in-process. Safe for concurrent use.
//
// Per-instance state (keys cache, revocation set) avoids the package-level
// global bug that arises when a single package hosts multiple validator instances.
type JWTValidator struct {
	jwksURL    string
	httpClient *http.Client

	// JWKS key cache
	keysMu   sync.RWMutex
	keys     map[string]*rsa.PublicKey
	keysTTL  time.Time
	cacheTTL time.Duration

	// Revocation set: sessionID → expiry time after which the entry can be dropped
	revokedMu sync.RWMutex
	revoked   map[string]time.Time
}

const defaultJWKSCacheTTL = 5 * time.Minute

// NewJWTValidator creates a validator that fetches JWKS from jwksURL.
// The JWKS response is cached for 5 minutes; a cache miss triggers one re-fetch.
func NewJWTValidator(jwksURL string) *JWTValidator {
	return &JWTValidator{
		jwksURL:    jwksURL,
		httpClient: &http.Client{Timeout: 15 * time.Second},
		keys:       make(map[string]*rsa.PublicKey),
		cacheTTL:   defaultJWKSCacheTTL,
		revoked:    make(map[string]time.Time),
	}
}

// ValidateToken verifies an RS256 JWT and returns its claims.
// Returns an error if the signature is invalid, the token is expired,
// or the session has been explicitly revoked via RevokeSession.
func (v *JWTValidator) ValidateToken(ctx context.Context, rawToken string) (*JWTClaims, error) {
	parts := strings.Split(rawToken, ".")
	if len(parts) != 3 {
		return nil, &jwtError{"malformed JWT"}
	}

	var header struct {
		Alg string `json:"alg"`
		Kid string `json:"kid"`
	}
	if err := decodeJWTSegment(parts[0], &header); err != nil {
		return nil, &jwtError{"invalid JWT header"}
	}
	if header.Alg != "RS256" {
		return nil, &jwtError{fmt.Sprintf("unsupported algorithm %q", header.Alg)}
	}

	key, err := v.getKey(ctx, header.Kid)
	if err != nil {
		return nil, &jwtError{err.Error()}
	}

	sigBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, &jwtError{"invalid JWT signature encoding"}
	}
	digest := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if err := rsa.VerifyPKCS1v15(key, crypto.SHA256, digest[:], sigBytes); err != nil {
		return nil, &jwtError{"invalid JWT signature"}
	}

	var claims struct {
		Sub         string   `json:"sub"`
		Email       string   `json:"email"`
		Exp         int64    `json:"exp"`
		Sid         string   `json:"sid"`
		Roles       []string `json:"roles"`
		RealmAccess struct {
			Roles []string `json:"roles"`
		} `json:"realm_access"`
	}
	if err := decodeJWTSegment(parts[1], &claims); err != nil {
		return nil, &jwtError{"invalid JWT claims"}
	}
	if claims.Sub == "" {
		return nil, &jwtError{"missing sub claim"}
	}
	if claims.Exp > 0 && time.Now().Unix() > claims.Exp {
		return nil, &jwtError{"token expired"}
	}

	// Session revocation check: if this session was explicitly revoked
	// (e.g. via admin logout) reject the token even though it's still
	// cryptographically valid. This closes the stateless-JWT revocation gap.
	if claims.Sid != "" && v.isRevoked(claims.Sid) {
		return nil, &jwtError{"session has been revoked"}
	}

	roles := append(claims.Roles, claims.RealmAccess.Roles...)
	return &JWTClaims{
		Subject:   claims.Sub,
		Email:     claims.Email,
		Roles:     roles,
		SessionID: claims.Sid,
		ExpiresAt: claims.Exp,
	}, nil
}

// RevokeSession marks sessionID as revoked until ttl elapses.
// Any subsequent ValidateToken call for a token carrying this session ID
// will return an error, regardless of the token's cryptographic validity.
//
// Set ttl to at least the maximum access-token lifetime for your IDP so that
// the revocation outlives any tokens that were issued for this session.
func (v *JWTValidator) RevokeSession(sessionID string, ttl time.Duration) {
	if sessionID == "" {
		return
	}
	expiry := time.Now().Add(ttl)

	v.revokedMu.Lock()
	v.revoked[sessionID] = expiry
	v.revokedMu.Unlock()
}

// isRevoked reports whether sessionID is in the revocation set.
// It prunes expired entries opportunistically.
func (v *JWTValidator) isRevoked(sessionID string) bool {
	now := time.Now()

	v.revokedMu.RLock()
	exp, found := v.revoked[sessionID]
	v.revokedMu.RUnlock()

	if !found {
		return false
	}
	if now.After(exp) {
		// Entry has expired — clean it up and report not revoked.
		v.revokedMu.Lock()
		delete(v.revoked, sessionID)
		v.revokedMu.Unlock()
		return false
	}
	return true
}

// getKey returns the RSA public key for the given kid, fetching JWKS if needed.
func (v *JWTValidator) getKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	v.keysMu.RLock()
	key, ok := v.keys[kid]
	fresh := time.Now().Before(v.keysTTL)
	v.keysMu.RUnlock()

	if ok && fresh {
		return key, nil
	}
	if err := v.fetchJWKS(ctx); err != nil {
		return nil, fmt.Errorf("fetch JWKS: %w", err)
	}
	v.keysMu.RLock()
	key, ok = v.keys[kid]
	v.keysMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown key id %q", kid)
	}
	return key, nil
}

// fetchJWKS fetches the JWKS from the configured URL and populates the key cache.
func (v *JWTValidator) fetchJWKS(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, v.jwksURL, nil)
	if err != nil {
		return err
	}
	resp, err := v.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var set struct {
		Keys []struct {
			Kid string `json:"kid"`
			Kty string `json:"kty"`
			Use string `json:"use"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&set); err != nil {
		return err
	}

	v.keysMu.Lock()
	defer v.keysMu.Unlock()
	for _, k := range set.Keys {
		if k.Kty != "RSA" || k.Use != "sig" {
			continue
		}
		nBytes, _ := base64.RawURLEncoding.DecodeString(k.N)
		eBytes, _ := base64.RawURLEncoding.DecodeString(k.E)
		if len(nBytes) == 0 || len(eBytes) == 0 {
			continue
		}
		v.keys[k.Kid] = &rsa.PublicKey{
			N: new(big.Int).SetBytes(nBytes),
			E: int(new(big.Int).SetBytes(eBytes).Int64()),
		}
	}
	v.keysTTL = time.Now().Add(v.cacheTTL)
	return nil
}

// decodeJWTSegment base64url-decodes a JWT segment and JSON-unmarshals it.
func decodeJWTSegment(seg string, v any) error {
	b, err := base64.RawURLEncoding.DecodeString(seg)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// jwtError is returned by ValidateToken for all validation failures.
// Callers can treat it as an "unauthorized" signal.
type jwtError struct{ msg string }

func (e *jwtError) Error() string { return e.msg }
