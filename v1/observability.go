package pluginsv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ── MonitoringFramework client ────────────────────────────────────────────────

type MonitoringFrameworkClient interface {
	IngestMetrics(ctx context.Context, in *IngestMetricsRequest, opts ...grpc.CallOption) (*IngestMetricsResponse, error)
	SupportsBillingMetrics(ctx context.Context, in *SupportsBillingMetricsRequest, opts ...grpc.CallOption) (*SupportsBillingMetricsResponse, error)
}

type monitoringFrameworkClient struct{ cc grpc.ClientConnInterface }

func NewMonitoringFrameworkClient(cc grpc.ClientConnInterface) MonitoringFrameworkClient {
	return &monitoringFrameworkClient{cc}
}

func (c *monitoringFrameworkClient) IngestMetrics(ctx context.Context, in *IngestMetricsRequest, opts ...grpc.CallOption) (*IngestMetricsResponse, error) {
	out := new(IngestMetricsResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.MonitoringFramework/IngestMetrics", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringFrameworkClient) SupportsBillingMetrics(ctx context.Context, in *SupportsBillingMetricsRequest, opts ...grpc.CallOption) (*SupportsBillingMetricsResponse, error) {
	out := new(SupportsBillingMetricsResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.MonitoringFramework/SupportsBillingMetrics", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── MonitoringFramework server ────────────────────────────────────────────────

type MonitoringFrameworkServer interface {
	IngestMetrics(context.Context, *IngestMetricsRequest) (*IngestMetricsResponse, error)
	SupportsBillingMetrics(context.Context, *SupportsBillingMetricsRequest) (*SupportsBillingMetricsResponse, error)
}

type UnimplementedMonitoringFrameworkServer struct{}

func (UnimplementedMonitoringFrameworkServer) IngestMetrics(context.Context, *IngestMetricsRequest) (*IngestMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestMetrics not implemented")
}

func (UnimplementedMonitoringFrameworkServer) SupportsBillingMetrics(context.Context, *SupportsBillingMetricsRequest) (*SupportsBillingMetricsResponse, error) {
	return &SupportsBillingMetricsResponse{Supported: false}, nil
}

func RegisterMonitoringFrameworkServer(s grpc.ServiceRegistrar, srv MonitoringFrameworkServer) {
	s.RegisterService(&MonitoringFramework_ServiceDesc, srv)
}

var MonitoringFramework_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.MonitoringFramework",
	HandlerType: (*MonitoringFrameworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestMetrics",
			Handler:    _MonitoringFramework_IngestMetrics_Handler,
		},
		{
			MethodName: "SupportsBillingMetrics",
			Handler:    _MonitoringFramework_SupportsBillingMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring.proto",
}

func _MonitoringFramework_IngestMetrics_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(IngestMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringFrameworkServer).IngestMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.MonitoringFramework/IngestMetrics"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MonitoringFrameworkServer).IngestMetrics(ctx, req.(*IngestMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringFramework_SupportsBillingMetrics_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SupportsBillingMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringFrameworkServer).SupportsBillingMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.MonitoringFramework/SupportsBillingMetrics"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MonitoringFrameworkServer).SupportsBillingMetrics(ctx, req.(*SupportsBillingMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
