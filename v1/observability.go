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

// ── MonitoringSource client ───────────────────────────────────────────────────

type MonitoringSourceClient interface {
	GetScrapeTargets(ctx context.Context, in *GetScrapeTargetsRequest, opts ...grpc.CallOption) (*GetScrapeTargetsResponse, error)
}

type monitoringSourceClient struct{ cc grpc.ClientConnInterface }

func NewMonitoringSourceClient(cc grpc.ClientConnInterface) MonitoringSourceClient {
	return &monitoringSourceClient{cc}
}

func (c *monitoringSourceClient) GetScrapeTargets(ctx context.Context, in *GetScrapeTargetsRequest, opts ...grpc.CallOption) (*GetScrapeTargetsResponse, error) {
	out := new(GetScrapeTargetsResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.MonitoringSource/GetScrapeTargets", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── MonitoringSource server ───────────────────────────────────────────────────

type MonitoringSourceServer interface {
	GetScrapeTargets(context.Context, *GetScrapeTargetsRequest) (*GetScrapeTargetsResponse, error)
}

type UnimplementedMonitoringSourceServer struct{}

func (UnimplementedMonitoringSourceServer) GetScrapeTargets(context.Context, *GetScrapeTargetsRequest) (*GetScrapeTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScrapeTargets not implemented")
}

func RegisterMonitoringSourceServer(s grpc.ServiceRegistrar, srv MonitoringSourceServer) {
	s.RegisterService(&MonitoringSource_ServiceDesc, srv)
}

var MonitoringSource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.MonitoringSource",
	HandlerType: (*MonitoringSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScrapeTargets",
			Handler:    _MonitoringSource_GetScrapeTargets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring.proto",
}

func _MonitoringSource_GetScrapeTargets_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetScrapeTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringSourceServer).GetScrapeTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.MonitoringSource/GetScrapeTargets"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MonitoringSourceServer).GetScrapeTargets(ctx, req.(*GetScrapeTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ── MonitoringLogs client ─────────────────────────────────────────────────────

type MonitoringLogsClient interface {
	IngestLog(ctx context.Context, in *IngestLogRequest, opts ...grpc.CallOption) (*IngestLogResponse, error)
}

type monitoringLogsClient struct{ cc grpc.ClientConnInterface }

func NewMonitoringLogsClient(cc grpc.ClientConnInterface) MonitoringLogsClient {
	return &monitoringLogsClient{cc}
}

func (c *monitoringLogsClient) IngestLog(ctx context.Context, in *IngestLogRequest, opts ...grpc.CallOption) (*IngestLogResponse, error) {
	out := new(IngestLogResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.MonitoringLogs/IngestLog", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── MonitoringLogs server ─────────────────────────────────────────────────────

type MonitoringLogsServer interface {
	IngestLog(context.Context, *IngestLogRequest) (*IngestLogResponse, error)
}

type UnimplementedMonitoringLogsServer struct{}

func (UnimplementedMonitoringLogsServer) IngestLog(context.Context, *IngestLogRequest) (*IngestLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestLog not implemented")
}

func RegisterMonitoringLogsServer(s grpc.ServiceRegistrar, srv MonitoringLogsServer) {
	s.RegisterService(&MonitoringLogs_ServiceDesc, srv)
}

var MonitoringLogs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.MonitoringLogs",
	HandlerType: (*MonitoringLogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestLog",
			Handler:    _MonitoringLogs_IngestLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring.proto",
}

func _MonitoringLogs_IngestLog_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(IngestLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringLogsServer).IngestLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.MonitoringLogs/IngestLog"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MonitoringLogsServer).IngestLog(ctx, req.(*IngestLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ── MonitoringTraces client ───────────────────────────────────────────────────

type MonitoringTracesClient interface {
	IngestSpan(ctx context.Context, in *IngestSpanRequest, opts ...grpc.CallOption) (*IngestSpanResponse, error)
}

type monitoringTracesClient struct{ cc grpc.ClientConnInterface }

func NewMonitoringTracesClient(cc grpc.ClientConnInterface) MonitoringTracesClient {
	return &monitoringTracesClient{cc}
}

func (c *monitoringTracesClient) IngestSpan(ctx context.Context, in *IngestSpanRequest, opts ...grpc.CallOption) (*IngestSpanResponse, error) {
	out := new(IngestSpanResponse)
	if err := c.cc.Invoke(ctx, "/kleff.plugins.v1.MonitoringTraces/IngestSpan", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ── MonitoringTraces server ───────────────────────────────────────────────────

type MonitoringTracesServer interface {
	IngestSpan(context.Context, *IngestSpanRequest) (*IngestSpanResponse, error)
}

type UnimplementedMonitoringTracesServer struct{}

func (UnimplementedMonitoringTracesServer) IngestSpan(context.Context, *IngestSpanRequest) (*IngestSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestSpan not implemented")
}

func RegisterMonitoringTracesServer(s grpc.ServiceRegistrar, srv MonitoringTracesServer) {
	s.RegisterService(&MonitoringTraces_ServiceDesc, srv)
}

var MonitoringTraces_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kleff.plugins.v1.MonitoringTraces",
	HandlerType: (*MonitoringTracesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestSpan",
			Handler:    _MonitoringTraces_IngestSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring.proto",
}

func _MonitoringTraces_IngestSpan_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(IngestSpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringTracesServer).IngestSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/kleff.plugins.v1.MonitoringTraces/IngestSpan"}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MonitoringTracesServer).IngestSpan(ctx, req.(*IngestSpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}
