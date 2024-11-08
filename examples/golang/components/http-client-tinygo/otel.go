package main

import (
	"net/http"
	"time"
	_ "unsafe"

	"github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/wasitel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func setupOTelSDK(client *http.Client) error {
	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// // Set up trace provider.
	// _, _ := newTraceProvider(client)
	tp, err := newTraceProvider(client)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tp)

	return nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(client *http.Client) (*trace.TracerProvider, error) {
	traceExporter, err := wasitel.New(wasitel.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}
