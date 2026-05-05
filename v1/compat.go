package pluginsv1

import "google.golang.org/grpc"

// Standard capability keys a plugin can declare via GetCapabilities.
const (
	CapabilityAPIMiddleware       = "api.middleware"
	CapabilityUIManifest          = "ui.manifest"
	CapabilityAPIRoutes           = "api.routes"
	CapabilityIdentityProvider    = "identity.provider"
	CapabilityIdentityFramework   = "identity.framework"
	CapabilityMonitoringMetrics   = "monitoring.metrics"
	CapabilityMonitoringLogs      = "monitoring.logs"
	CapabilityMonitoringTraces    = "monitoring.traces"
	CapabilityMonitoringSource    = "monitoring.source"
)

// ── Backward-compatible client type aliases ──────────────────────────────────
// These names were used before the proto-gen overhaul renamed the services.

type PluginMiddlewareClient = APIMiddlewareClient
type PluginUIClient = UIManifestServiceClient
type PluginHTTPClient = APIRoutesClient

// HandleHTTPRequest and HandleHTTPResponse are the old names for the route
// handler request/response types, now generated as HandleRequest/HandleResponse.
type HandleHTTPRequest = HandleRequest
type HandleHTTPResponse = HandleResponse

// NewPluginMiddlewareClient, NewPluginUIClient, NewPluginHTTPClient are the old
// constructor names; they forward to the generated constructors.
func NewPluginMiddlewareClient(cc grpc.ClientConnInterface) PluginMiddlewareClient {
	return NewAPIMiddlewareClient(cc)
}
func NewPluginUIClient(cc grpc.ClientConnInterface) PluginUIClient {
	return NewUIManifestServiceClient(cc)
}
func NewPluginHTTPClient(cc grpc.ClientConnInterface) PluginHTTPClient {
	return NewAPIRoutesClient(cc)
}

// Layer tag declarations for plugin manifests.
const (
	TagFrontend = "frontend"
	TagBackend  = "backend"
	TagIdentity = "identity"
	TagDevOps   = "devops"
)

// HealthStatus aliases the generated HealthResponse_Status enum for backward compatibility.
type HealthStatus = HealthResponse_Status

// HealthStatus constants matching the generated HealthResponse_Status enum values.
const (
	HealthStatusUnknown  HealthStatus = HealthResponse_UNKNOWN
	HealthStatusHealthy  HealthStatus = HealthResponse_HEALTHY
	HealthStatusDegraded HealthStatus = HealthResponse_DEGRADED
	HealthStatusUnhealthy HealthStatus = HealthResponse_UNHEALTHY
)

// ErrorCode aliases the generated Error_Code enum for backward compatibility.
type ErrorCode = Error_Code

// ErrorCode constants matching the generated Error_Code enum values.
const (
	ErrorCodeUnknown         ErrorCode = Error_UNKNOWN
	ErrorCodeInvalidArgument ErrorCode = Error_INVALID_ARGUMENT
	ErrorCodeUnauthorized    ErrorCode = Error_UNAUTHORIZED
	ErrorCodeConflict        ErrorCode = Error_CONFLICT
	ErrorCodeNotFound        ErrorCode = Error_NOT_FOUND
	ErrorCodeInternal        ErrorCode = Error_INTERNAL
	ErrorCodeNotSupported    ErrorCode = Error_NOT_SUPPORTED
)

// PluginError is a backward-compatible Go error type that wraps a proto Error.
// Callers that previously type-asserted errors to *PluginError can continue to do so.
type PluginError struct {
	Code    ErrorCode
	Message string
}

func (e *PluginError) Error() string { return e.Message }

// ProtoErrorToPluginError converts a proto *Error into a *PluginError.
// Returns nil if e is nil.
func ProtoErrorToPluginError(e *Error) *PluginError {
	if e == nil {
		return nil
	}
	return &PluginError{Code: e.Code, Message: e.Message}
}
