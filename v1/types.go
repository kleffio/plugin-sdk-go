package pluginsv1

// ── Health types ──────────────────────────────────────────────────────────────

type HealthRequest struct{}

type HealthResponse struct {
	Status  HealthStatus `json:"status"`
	Message string       `json:"message,omitempty"`
}

type HealthStatus int32

const (
	HealthStatusUnknown   HealthStatus = 0
	HealthStatusHealthy   HealthStatus = 1
	HealthStatusDegraded  HealthStatus = 2
	HealthStatusUnhealthy HealthStatus = 3
)

// ── Error type ────────────────────────────────────────────────────────────────

type PluginError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message,omitempty"`
}

func (e *PluginError) Error() string { return e.Message }

type ErrorCode int32

const (
	ErrorCodeUnknown         ErrorCode = 0
	ErrorCodeInvalidArgument ErrorCode = 1
	ErrorCodeUnauthorized    ErrorCode = 2
	ErrorCodeConflict        ErrorCode = 3
	ErrorCodeNotFound        ErrorCode = 4
	ErrorCodeInternal        ErrorCode = 5
	ErrorCodeNotSupported    ErrorCode = 6
)

// ── Token validation types ────────────────────────────────────────────────────

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	Claims *TokenClaims `json:"claims,omitempty"`
	Error  *PluginError `json:"error,omitempty"`
}

// TokenClaims carries the verified identity extracted from a token.
type TokenClaims struct {
	Subject  string   `json:"subject"`
	Username string   `json:"username,omitempty"`
	Email    string   `json:"email,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

// ── OIDC config types ─────────────────────────────────────────────────────────

type GetOIDCConfigRequest struct{}

// OIDCConfig carries the OIDC parameters the frontend needs to initialise its
// OIDC client library.
type OIDCConfig struct {
	Authority string   `json:"authority"`
	ClientID  string   `json:"client_id"`
	JwksURI   string   `json:"jwks_uri"`
	Scopes    []string `json:"scopes,omitempty"`
	// AuthMode controls whether the panel uses a headless login form or an
	// OIDC Authorization Code redirect to the IDP. Valid values: "headless"
	// (default) or "redirect".
	AuthMode string `json:"auth_mode,omitempty"`
	// Pre-fetched OIDC endpoint URLs. When set, the frontend passes them as
	// metadata to oidc-client-ts so it skips the cross-origin discovery fetch.
	TokenEndpoint         string `json:"token_endpoint,omitempty"`
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
	UserinfoEndpoint      string `json:"userinfo_endpoint,omitempty"`
	EndSessionEndpoint    string `json:"end_session_endpoint,omitempty"`
	// InternalTokenEndpoint is the Docker-internal token endpoint URL.
	// The platform uses it to proxy PKCE token exchanges server-side so the
	// browser never has to POST cross-origin to the IDP.
	InternalTokenEndpoint string `json:"internal_token_endpoint,omitempty"`
}

type GetOIDCConfigResponse struct {
	Config *OIDCConfig  `json:"config,omitempty"`
	Error  *PluginError `json:"error,omitempty"`
}

// ── IDP types ─────────────────────────────────────────────────────────────────

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token *TokenSet    `json:"token,omitempty"`
	Error *PluginError `json:"error,omitempty"`
}

type TokenSet struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope,omitempty"`
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type RegisterResponse struct {
	UserID string       `json:"user_id,omitempty"`
	Error  *PluginError `json:"error,omitempty"`
}

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User  *UserInfo    `json:"user,omitempty"`
	Error *PluginError `json:"error,omitempty"`
}

type UserInfo struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// ── Layer tag declarations ────────────────────────────────────────────────────

// Layer tags declare which part of the stack a plugin touches.
// The platform uses these to enforce capability permissions — a plugin may only
// exercise capabilities that its declared layer tags permit.
const (
	// TagFrontend allows ui.manifest. Use for plugins that only contribute UI.
	TagFrontend = "frontend"

	// TagBackend allows api.middleware and api.routes. Use for plugins that
	// intercept or own HTTP routes on the platform API.
	TagBackend = "backend"

	// TagIdentity allows identity.provider. Use for authentication/IDP plugins.
	TagIdentity = "identity"

	// TagDevOps is reserved for future daemon and Kubernetes capabilities.
	TagDevOps = "devops"
)

// ── Capability declarations ───────────────────────────────────────────────────

// Standard capability keys a plugin can declare via GetCapabilities.
const (
	// CapabilityAPIMiddleware: plugin implements PluginMiddleware.OnRequest.
	// The platform calls it before every authenticated handler.
	CapabilityAPIMiddleware = "api.middleware"

	// CapabilityUIManifest: plugin implements PluginUI.GetUIManifest.
	// The panel fetches aggregated manifests to render plugin-contributed UI.
	CapabilityUIManifest = "ui.manifest"

	// CapabilityAPIRoutes: plugin implements PluginHTTP.GetRoutes + Handle.
	// The platform intercepts declared routes and forwards them to the plugin
	// as raw HTTP requests. The plugin owns 100% of the response.
	CapabilityAPIRoutes = "api.routes"

	// CapabilityIdentityProvider: plugin is an authentication/identity provider.
	// The platform uses this to determine whether auth is available, exposing
	// GET /api/v1/auth/config with enabled:false when no such plugin is active.
	CapabilityIdentityProvider = "identity.provider"

	// CapabilityMonitoringFramework: plugin is a monitoring framework (time-series
	// storage + query). The platform forwards workload metric samples to it via IngestMetrics.
	CapabilityMonitoringFramework = "monitoring.framework"

	// CapabilityMonitoringProvider: plugin is a metrics collector (e.g. cAdvisor,
	// node_exporter, Promtail). Reserved for future use.
	CapabilityMonitoringProvider = "monitoring.provider"
)

type GetCapabilitiesRequest struct{}

type GetCapabilitiesResponse struct {
	Capabilities []string `json:"capabilities"`
}

// ── API middleware types ───────────────────────────────────────────────────────

// MiddlewareRequest is sent to PluginMiddleware.OnRequest for every authenticated
// HTTP request. The platform has already validated the token; UserID and Roles
// are the verified identity. The plugin may allow or deny the request.
type MiddlewareRequest struct {
	UserID  string            `json:"user_id"`
	Roles   []string          `json:"roles,omitempty"`
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers,omitempty"`
}

// MiddlewareResponse is returned by PluginMiddleware.OnRequest.
// If Allow is false, the platform rejects the request with the Error message.
type MiddlewareResponse struct {
	Allow bool         `json:"allow"`
	Error *PluginError `json:"error,omitempty"`
}

// ── UI manifest types ─────────────────────────────────────────────────────────

// NavItem represents a navigation item contributed by a plugin.
type NavItem struct {
	Label      string     `json:"label"`
	Icon       string     `json:"icon,omitempty"`       // Lucide icon name
	Path       string     `json:"path"`                 // panel route path
	Permission string     `json:"permission,omitempty"` // required role; empty = all authenticated users
	Children   []*NavItem `json:"children,omitempty"`
}

// SettingsPage represents a settings page contributed by a plugin.
// If IframeURL is set the panel embeds it; otherwise it renders an empty shell.
type SettingsPage struct {
	Label     string `json:"label"`
	Path      string `json:"path"`                // panel route, e.g. /settings/keycloak
	IframeURL string `json:"iframe_url,omitempty"` // if set, embed as iframe
}

// LoginConfig lets an IDP plugin customise the panel's headless login form.
type LoginConfig struct {
	// DisableSignupLink hides the "Create account" link on the login page.
	DisableSignupLink bool `json:"disable_signup_link,omitempty"`
}

// SignupConfig lets an IDP plugin customise the self-registration page.
type SignupConfig struct {
	// Disabled replaces the signup form with a message directing the user to
	// register through the identity provider directly.
	Disabled bool `json:"disabled,omitempty"`
	// HideFirstName omits the first-name field.
	HideFirstName bool `json:"hide_first_name,omitempty"`
	// HideLastName omits the last-name field.
	HideLastName bool `json:"hide_last_name,omitempty"`
	// HideUsername omits the username field.
	HideUsername bool `json:"hide_username,omitempty"`
}

// ProfileSection is a plugin-contributed section rendered on the user settings page.
type ProfileSection struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	// IframeURL, if set, embeds the URL in an iframe inside the section card.
	IframeURL string `json:"iframe_url,omitempty"`
	// Actions lists the native actions this section supports (e.g. "change_password").
	// The panel renders built-in Kleff-styled forms for each declared action.
	Actions []string `json:"actions,omitempty"`
}

// ── Session types ─────────────────────────────────────────────────────────────

// Session represents one active login session for a user.
type Session struct {
	ID          string `json:"id"`
	IPAddress   string `json:"ip_address,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
	StartedAt   int64  `json:"started_at,omitempty"` // unix seconds
	LastAccess  int64  `json:"last_access,omitempty"` // unix seconds
	Current     bool   `json:"current"`
}

type ListSessionsRequest struct {
	UserID           string `json:"user_id"`
	CurrentSessionID string `json:"current_session_id,omitempty"` // hints which session to mark current
}

type ListSessionsResponse struct {
	Sessions []*Session   `json:"sessions,omitempty"`
	Error    *PluginError `json:"error,omitempty"`
}

type RevokeSessionRequest struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type RevokeSessionResponse struct {
	Error *PluginError `json:"error,omitempty"`
}

// ── Change password types ─────────────────────────────────────────────────────

// ChangePasswordRequest is sent by the platform to change a user's password.
// UserID is the IDP subject claim. CurrentPassword is verified by the plugin
// before the new password is applied.
type ChangePasswordRequest struct {
	UserID          string `json:"user_id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Error *PluginError `json:"error,omitempty"`
}

// UIManifest aggregates all UI contributions from a single plugin.
type UIManifest struct {
	PluginID        string            `json:"plugin_id"`
	NavItems        []*NavItem        `json:"nav_items,omitempty"`
	SettingsPages   []*SettingsPage   `json:"settings_pages,omitempty"`
	LoginConfig     *LoginConfig      `json:"login_config,omitempty"`
	SignupConfig    *SignupConfig     `json:"signup_config,omitempty"`
	ProfileSections []*ProfileSection `json:"profile_sections,omitempty"`
}

type GetUIManifestRequest struct{}

type GetUIManifestResponse struct {
	Manifest *UIManifest  `json:"manifest,omitempty"`
	Error    *PluginError `json:"error,omitempty"`
}

// ── HTTP route ownership types ─────────────────────────────────────────────────

// Route declares one HTTP endpoint the plugin wants to own.
// The platform will intercept matching requests and forward them via Handle.
type Route struct {
	Method string `json:"method"` // HTTP method ("GET", "POST", …) or "*" for any
	Path   string `json:"path"`   // exact path or prefix ending in "*" (e.g. "/api/v1/auth/*")
	Public bool   `json:"public"` // reserved for future use; all auth is handled by plugin middleware
}

type GetRoutesRequest struct{}

type GetRoutesResponse struct {
	Routes []*Route     `json:"routes"`
	Error  *PluginError `json:"error,omitempty"`
}

// HTTPRequest is the forwarded HTTP request the plugin receives via Handle.
// UserID and Roles are injected by the platform after token validation for non-public routes.
type HTTPRequest struct {
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	RawQuery string            `json:"raw_query,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
	Body     []byte            `json:"body,omitempty"`
	UserID   string            `json:"user_id,omitempty"`
	Roles    []string          `json:"roles,omitempty"`
}

// HTTPResponse is what the plugin returns; the platform writes it verbatim.
type HTTPResponse struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       []byte            `json:"body"`
}

type HandleHTTPRequest struct {
	Request *HTTPRequest `json:"request"`
}

type HandleHTTPResponse struct {
	Response *HTTPResponse `json:"response,omitempty"`
	Error    *PluginError  `json:"error,omitempty"`
}

// ── Token refresh types ───────────────────────────────────────────────────────

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token *TokenSet    `json:"token,omitempty"`
	Error *PluginError `json:"error,omitempty"`
}

// ── Observability types ───────────────────────────────────────────────────────

// MetricSample carries a single resource-usage observation for one workload.
type MetricSample struct {
	WorkloadID    string  `json:"workload_id"`
	NodeID        string  `json:"node_id"`
	OrgID         string  `json:"org_id"`
	ProjectID     string  `json:"project_id"`
	Timestamp     int64   `json:"timestamp"` // unix seconds
	CPUMillicores int64   `json:"cpu_millicores"`
	MemoryMB      int64   `json:"memory_mb"`
	NetworkRxMB   float64 `json:"network_rx_mb"`
	NetworkTxMB   float64 `json:"network_tx_mb"`
	DiskReadMB    float64 `json:"disk_read_mb"`
	DiskWriteMB   float64 `json:"disk_write_mb"`
}

type IngestMetricsRequest struct {
	Sample *MetricSample `json:"sample"`
}

type IngestMetricsResponse struct {
	Error *PluginError `json:"error,omitempty"`
}

type SupportsBillingMetricsRequest struct{}

type SupportsBillingMetricsResponse struct {
	Supported bool `json:"supported"`
}

// ── Admin seeding types ───────────────────────────────────────────────────────

// EnsureAdminRequest is sent by the platform after installing an IDP plugin to
// trigger admin-user seeding. Each IDP plugin implements this according to its
// own system (Keycloak realm roles, Authentik groups, etc.).
type EnsureAdminRequest struct{}

type EnsureAdminResponse struct {
	Error *PluginError `json:"error,omitempty"`
}

