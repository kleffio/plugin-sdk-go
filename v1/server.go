package pluginsv1

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// BasePlugin provides the boilerplate every plugin needs: health/capabilities
// server, gRPC server setup, and graceful shutdown. Embed it in your plugin
// struct and override the capability methods you implement.
//
//	type MyPlugin struct {
//	    pluginsv1.BasePlugin
//	}
//
// Call Serve() from main() to start the gRPC server and block until SIGTERM.
type BasePlugin struct {
	UnimplementedPluginHealthServer

	name         string
	version      string
	capabilities []string
	logger       *slog.Logger
}

// NewBasePlugin creates a BasePlugin. name and version are reported by the
// Health RPC. capabilities lists the capability strings the plugin declares
// (e.g. CapabilityUIManifest, CapabilityIdentityProvider).
func NewBasePlugin(name, version string, capabilities []string) *BasePlugin {
	return &BasePlugin{
		name:         name,
		version:      version,
		capabilities: capabilities,
		logger:       slog.Default(),
	}
}

// WithLogger overrides the logger used for startup/shutdown messages.
func (b *BasePlugin) WithLogger(l *slog.Logger) *BasePlugin {
	b.logger = l
	return b
}

// Health implements PluginHealthServer.
func (b *BasePlugin) Health(_ context.Context, _ *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{Status: HealthStatusHealthy, Message: b.name + " " + b.version}, nil
}

// GetCapabilities implements PluginHealthServer.
func (b *BasePlugin) GetCapabilities(_ context.Context, _ *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error) {
	return &GetCapabilitiesResponse{Capabilities: b.capabilities}, nil
}

// Serve starts the gRPC server on the given address, registers all provided
// service registrars, and blocks until SIGTERM/SIGINT. It always registers
// the PluginHealth service (backed by the BasePlugin itself).
//
// Example:
//
//	p := &MyPlugin{BasePlugin: *pluginsv1.NewBasePlugin("my-plugin", "1.0.0", []string{pluginsv1.CapabilityUIManifest})}
//	pluginsv1.Serve(":50051", p, func(s *grpc.Server) {
//	    pluginsv1.RegisterMyPluginServer(s, p)
//	})
func Serve(addr string, base *BasePlugin, register ...func(*grpc.Server)) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("plugin serve: listen %s: %w", addr, err)
	}

	srv := grpc.NewServer()
	RegisterPluginHealthServer(srv, base)
	for _, fn := range register {
		fn(srv)
	}

	base.logger.Info("plugin gRPC server starting", "addr", addr, "plugin", base.name, "version", base.version)

	errCh := make(chan error, 1)
	go func() { errCh <- srv.Serve(lis) }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errCh:
		return fmt.Errorf("plugin serve: %w", err)
	case sig := <-quit:
		base.logger.Info("plugin gRPC server shutting down", "signal", sig)
		srv.GracefulStop()
		return nil
	}
}

// DefaultAddr returns the standard plugin gRPC address from the PORT env var,
// defaulting to :50051 if unset.
func DefaultAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":50051"
}
