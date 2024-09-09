package otlp

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func SpanExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create otlp exporter: %w", err)
	}
	return exporter, nil
}
