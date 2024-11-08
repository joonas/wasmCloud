package trace

type ExportTraceServiceRequest struct {
	ResourceSpanss []*ResourceSpans `json:"resource_spans,omitempty"`
}
