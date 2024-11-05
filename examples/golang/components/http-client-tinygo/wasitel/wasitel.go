package wasitel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/sdk/trace"
)

var _ trace.SpanExporter = &Exporter{}

const defaultUrl = "http://localhost:9999"

// TODO: Consider adding context as the first argument here
func New(opts ...Option) (*Exporter, error) {
	cfg := newConfig(opts...)
	exporter := &Exporter{
		cfg.HttpClient,
	}
	return exporter, nil
}

type Exporter struct {
	client *http.Client
}

func (e *Exporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	// TODO: Check if the exporter is shutting down before attempting to export
	if len(spans) == 0 {
		// Nothing to export
		return nil
	}
	resourceSpans := convertSpans(spans)
	body, err := json.Marshal(resourceSpans)
	if err != nil {
		return fmt.Errorf("failed to serialize zipkin models to JSON: %w", err)
	}
	// req, err := http.NewRequestWithContext(ctx, http.MethodPost, defaultUrl, bytes.NewBuffer(body))
	req, err := http.NewRequest(http.MethodPost, defaultUrl, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request to %s: %w", defaultUrl, err)
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = e.client.Do(req)
	if err != nil {
		return fmt.Errorf("request to %s failed: %w", defaultUrl, err)
	}

	return nil
}

func (e *Exporter) Shutdown(ctx context.Context) error {
	return nil
}
