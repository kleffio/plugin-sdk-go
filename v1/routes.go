package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── PluginHTTP client ─────────────────────────────────────────────────────────

// PluginHTTPClient is implemented by plugins that declare CapabilityAPIRoutes.
// The platform calls GetRoutes once on connect, then Handle for every matched request.
type PluginHTTPClient interface {
	GetRoutes(ctx context.Context, in *GetRoutesRequest, opts ...grpc.CallOption) (*GetRoutesResponse, error)
	Handle(ctx context.Context, in *HandleHTTPRequest, opts ...grpc.CallOption) (*HandleHTTPResponse, error)
}

type pluginHTTPClient struct{ cc grpc.ClientConnInterface }

func NewPluginHTTPClient(cc grpc.ClientConnInterface) PluginHTTPClient {
	return &pluginHTTPClient{cc}
}

func (c *pluginHTTPClient) GetRoutes(ctx context.Context, in *GetRoutesRequest, opts ...grpc.CallOption) (*GetRoutesResponse, error) {
	out := new(GetRoutesResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginHTTP/GetRoutes", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginHTTPClient) Handle(ctx context.Context, in *HandleHTTPRequest, opts ...grpc.CallOption) (*HandleHTTPResponse, error) {
	out := new(HandleHTTPResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginHTTP/Handle", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── PluginHTTP server ─────────────────────────────────────────────────────────

// PluginHTTPServer is the server interface plugins implement to own HTTP routes.
type PluginHTTPServer interface {
	GetRoutes(context.Context, *GetRoutesRequest) (*GetRoutesResponse, error)
	Handle(context.Context, *HandleHTTPRequest) (*HandleHTTPResponse, error)
}

type UnimplementedPluginHTTPServer struct{}

func (UnimplementedPluginHTTPServer) GetRoutes(context.Context, *GetRoutesRequest) (*GetRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoutes not implemented")
}

func (UnimplementedPluginHTTPServer) Handle(context.Context, *HandleHTTPRequest) (*HandleHTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterPluginHTTPServer(s grpc.ServiceRegistrar, srv PluginHTTPServer) {
	s.RegisterService(&PluginHTTP_ServiceDesc, srv)
}

var PluginHTTP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.PluginHTTP",
	HandlerType: (*PluginHTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoutes",
			Handler:    _PluginHTTP_GetRoutes_Handler,
		},
		{
			MethodName: "Handle",
			Handler:    _PluginHTTP_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pluginhttp.proto",
}

func _PluginHTTP_GetRoutes_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginHTTPServer).GetRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginHTTP/GetRoutes"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginHTTPServer).GetRoutes(ctx, req.(*GetRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginHTTP_Handle_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(HandleHTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginHTTPServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginHTTP/Handle"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginHTTPServer).Handle(ctx, req.(*HandleHTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}
