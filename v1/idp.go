package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── IdentityPlugin client ─────────────────────────────────────────────────────

type IdentityPluginClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// ValidateToken verifies a bearer token and returns its claims.
	// The platform calls this on every authenticated request.
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	// GetOIDCConfig returns the OIDC configuration for the frontend.
	GetOIDCConfig(ctx context.Context, in *GetOIDCConfigRequest, opts ...grpc.CallOption) (*GetOIDCConfigResponse, error)
	// RefreshToken exchanges a refresh token for a new token set.
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
}

type identityPluginClient struct{ cc grpc.ClientConnInterface }

func NewIdentityPluginClient(cc grpc.ClientConnInterface) IdentityPluginClient {
	return &identityPluginClient{cc}
}

func (c *identityPluginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/Login", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/Register", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/GetUser", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/Health", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/ValidateToken", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) GetOIDCConfig(ctx context.Context, in *GetOIDCConfigRequest, opts ...grpc.CallOption) (*GetOIDCConfigResponse, error) {
	out := new(GetOIDCConfigResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/GetOIDCConfig", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityPluginClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.IdentityPlugin/RefreshToken", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── IdentityPlugin server ─────────────────────────────────────────────────────

type IdentityPluginServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	GetOIDCConfig(context.Context, *GetOIDCConfigRequest) (*GetOIDCConfigResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
}

type UnimplementedIdentityPluginServer struct{}

func (UnimplementedIdentityPluginServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIdentityPluginServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIdentityPluginServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIdentityPluginServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedIdentityPluginServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedIdentityPluginServer) GetOIDCConfig(context.Context, *GetOIDCConfigRequest) (*GetOIDCConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOIDCConfig not implemented")
}
func (UnimplementedIdentityPluginServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

func RegisterIdentityPluginServer(s grpc.ServiceRegistrar, srv IdentityPluginServer) {
	s.RegisterService(&IdentityPlugin_ServiceDesc, srv)
}

var IdentityPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.IdentityPlugin",
	HandlerType: (*IdentityPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Login", Handler: _IdentityPlugin_Login_Handler},
		{MethodName: "Register", Handler: _IdentityPlugin_Register_Handler},
		{MethodName: "GetUser", Handler: _IdentityPlugin_GetUser_Handler},
		{MethodName: "Health", Handler: _IdentityPlugin_Health_Handler},
		{MethodName: "ValidateToken", Handler: _IdentityPlugin_ValidateToken_Handler},
		{MethodName: "GetOIDCConfig", Handler: _IdentityPlugin_GetOIDCConfig_Handler},
		{MethodName: "RefreshToken", Handler: _IdentityPlugin_RefreshToken_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idp.proto",
}

func _IdentityPlugin_Login_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/Login"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_Register_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/Register"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_GetUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/GetUser"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_Health_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/Health"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_ValidateToken_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/ValidateToken"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_GetOIDCConfig_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetOIDCConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).GetOIDCConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/GetOIDCConfig"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).GetOIDCConfig(ctx, req.(*GetOIDCConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityPlugin_RefreshToken_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityPluginServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.IdentityPlugin/RefreshToken"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(IdentityPluginServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}
