package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── PluginUI client ───────────────────────────────────────────────────────────

// PluginUIClient is implemented by plugins that declare CapabilityUIManifest.
// The platform calls GetUIManifest to collect nav items and settings pages to
// surface in the panel.
type PluginUIClient interface {
	GetUIManifest(ctx context.Context, in *GetUIManifestRequest, opts ...grpc.CallOption) (*GetUIManifestResponse, error)
}

type pluginUIClient struct{ cc grpc.ClientConnInterface }

func NewPluginUIClient(cc grpc.ClientConnInterface) PluginUIClient {
	return &pluginUIClient{cc}
}

func (c *pluginUIClient) GetUIManifest(ctx context.Context, in *GetUIManifestRequest, opts ...grpc.CallOption) (*GetUIManifestResponse, error) {
	out := new(GetUIManifestResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.PluginUI/GetUIManifest", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── PluginUI server ───────────────────────────────────────────────────────────

// PluginUIServer is the server interface plugins implement to contribute UI.
type PluginUIServer interface {
	GetUIManifest(context.Context, *GetUIManifestRequest) (*GetUIManifestResponse, error)
}

type UnimplementedPluginUIServer struct{}

func (UnimplementedPluginUIServer) GetUIManifest(context.Context, *GetUIManifestRequest) (*GetUIManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUIManifest not implemented")
}

func RegisterPluginUIServer(s grpc.ServiceRegistrar, srv PluginUIServer) {
	s.RegisterService(&PluginUI_ServiceDesc, srv)
}

var PluginUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.PluginUI",
	HandlerType: (*PluginUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUIManifest",
			Handler:    _PluginUI_GetUIManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ui.proto",
}

func _PluginUI_GetUIManifest_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetUIManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginUIServer).GetUIManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.PluginUI/GetUIManifest"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PluginUIServer).GetUIManifest(ctx, req.(*GetUIManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}
