package jaeger

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

var jaegerEndpoint string

func init() {
	jaegerEndpoint = os.Getenv("JAEGER_ENDPOINT")
	if jaegerEndpoint == "" {
		jaegerEndpoint = "http://localhost:14268/api/traces" // default value
	}
}

func SpanExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	traceExporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	return traceExporter, nil
}
