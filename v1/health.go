package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── PluginHealth client ───────────────────────────────────────────────────────

type PluginHealthClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	GetCapabilities(ctx context.Context, in *GetCapabilitiesRequest, opts ...grpc.CallOption) (*GetCapabilitiesResponse, error)
}

type pluginHealthClient struct{ cc grpc.ClientConnInterface }

func NewPluginHealthClient(cc grpc.ClientConnInterface) PluginHealthClient {
	return &pluginHealthClient{cc}
}

func (c *pluginHealthClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginHealth/Health", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginHealthClient) GetCapabilities(ctx context.Context, in *GetCapabilitiesRequest, opts ...grpc.CallOption) (*GetCapabilitiesResponse, error) {
	out := new(GetCapabilitiesResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginHealth/GetCapabilities", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── PluginHealth server ───────────────────────────────────────────────────────

type PluginHealthServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	// GetCapabilities declares which optional extension points this plugin implements.
	// Return an empty list (not an error) if the plugin has no extra capabilities.
	GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error)
}

type UnimplementedPluginHealthServer struct{}

func (UnimplementedPluginHealthServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func (UnimplementedPluginHealthServer) GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error) {
	return &GetCapabilitiesResponse{}, nil // no capabilities by default
}

func RegisterPluginHealthServer(s grpc.ServiceRegistrar, srv PluginHealthServer) {
	s.RegisterService(&PluginHealth_ServiceDesc, srv)
}

var PluginHealth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.PluginHealth",
	HandlerType: (*PluginHealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _PluginHealth_Health_Handler,
		},
		{
			MethodName: "GetCapabilities",
			Handler:    _PluginHealth_GetCapabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}

func _PluginHealth_Health_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginHealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginHealth/Health"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginHealthServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginHealth_GetCapabilities_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginHealthServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginHealth/GetCapabilities"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginHealthServer).GetCapabilities(ctx, req.(*GetCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}
