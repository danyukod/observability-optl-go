package stdout

import (
	"fmt"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func SpanExporter() (sdktrace.SpanExporter, error) {
	exporter, err := stdouttrace.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout exporter: %w", err)
	}
	return exporter, nil
}
