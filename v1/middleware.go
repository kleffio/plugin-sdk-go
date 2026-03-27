package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── PluginMiddleware client ───────────────────────────────────────────────────

// PluginMiddlewareClient is implemented by plugins that declare CapabilityAPIMiddleware.
// The platform calls OnRequest before every authenticated HTTP handler.
type PluginMiddlewareClient interface {
	OnRequest(ctx context.Context, in *MiddlewareRequest, opts ...grpc.CallOption) (*MiddlewareResponse, error)
}

type pluginMiddlewareClient struct{ cc grpc.ClientConnInterface }

func NewPluginMiddlewareClient(cc grpc.ClientConnInterface) PluginMiddlewareClient {
	return &pluginMiddlewareClient{cc}
}

func (c *pluginMiddlewareClient) OnRequest(ctx context.Context, in *MiddlewareRequest, opts ...grpc.CallOption) (*MiddlewareResponse, error) {
	out := new(MiddlewareResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginMiddleware/OnRequest", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── PluginMiddleware server ───────────────────────────────────────────────────

// PluginMiddlewareServer is the server interface plugins implement to intercept requests.
type PluginMiddlewareServer interface {
	OnRequest(context.Context, *MiddlewareRequest) (*MiddlewareResponse, error)
}

type UnimplementedPluginMiddlewareServer struct{}

func (UnimplementedPluginMiddlewareServer) OnRequest(context.Context, *MiddlewareRequest) (*MiddlewareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnRequest not implemented")
}

func RegisterPluginMiddlewareServer(s grpc.ServiceRegistrar, srv PluginMiddlewareServer) {
	s.RegisterService(&PluginMiddleware_ServiceDesc, srv)
}

var PluginMiddleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.PluginMiddleware",
	HandlerType: (*PluginMiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnRequest",
			Handler:    _PluginMiddleware_OnRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middleware.proto",
}

func _PluginMiddleware_OnRequest_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(MiddlewareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginMiddlewareServer).OnRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginMiddleware/OnRequest"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginMiddlewareServer).OnRequest(ctx, req.(*MiddlewareRequest))
	}
	return interceptor(ctx, in, info, handler)
}
