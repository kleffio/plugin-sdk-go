package pluginsv1

// MetricSample is a single resource-usage snapshot for a workload.
type MetricSample struct {
	Timestamp    int64
	WorkloadID   string
	WorkloadName string
	NodeID       string
	OrgID        string
	ProjectID    string
	CPUMillicores int64
	MemoryMB      int64
	NetworkRxMB   float64
	NetworkTxMB   float64
	DiskReadMB    float64
	DiskWriteMB   float64
}

// IngestMetricsRequest carries a single metric sample to the plugin.
type IngestMetricsRequest struct {
	Sample *MetricSample
}

// IngestMetricsResponse is returned by IngestMetrics.
type IngestMetricsResponse struct {
	Error *PluginError
}

// SupportsBillingMetricsRequest is the request for SupportsBillingMetrics.
type SupportsBillingMetricsRequest struct{}

// SupportsBillingMetricsResponse indicates whether billing-grade metrics are available.
type SupportsBillingMetricsResponse struct {
	Supported bool
}

// ScrapeTarget describes a Prometheus-compatible scrape endpoint.
type ScrapeTarget struct {
	Address string
	Labels  map[string]string
}

// GetScrapeTargetsRequest is the request for GetScrapeTargets.
type GetScrapeTargetsRequest struct{}

// GetScrapeTargetsResponse carries the list of scrape targets.
type GetScrapeTargetsResponse struct {
	Targets []*ScrapeTarget
}

// LogEntry is a single structured log line from a workload.
type LogEntry struct {
	WorkloadID   string
	WorkloadName string
	OrgID        string
	ProjectID    string
	Level        string
	Labels       map[string]string
	Timestamp    int64
	Message      string
}

// IngestLogRequest carries a single log entry to the plugin.
type IngestLogRequest struct {
	Entry *LogEntry
}

// IngestLogResponse is returned by IngestLog.
type IngestLogResponse struct {
	Error *PluginError
}

// Span is a single distributed-tracing span from a workload.
type Span struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
	Name         string
	StartTimeNs  int64
	EndTimeNs    int64
	WorkloadID   string
	OrgID        string
	ProjectID    string
	Status       string
	Attributes   map[string]string
}

// IngestSpanRequest carries a single span to the plugin.
type IngestSpanRequest struct {
	Span *Span
}

// IngestSpanResponse is returned by IngestSpan.
type IngestSpanResponse struct {
	Error *PluginError
}
