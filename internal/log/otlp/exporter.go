package otlp

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
)

func Exporter(ctx context.Context) (log.Exporter, error) {
	exporter, err := otlploghttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create otlplog exporter: %w", err)
	}
	return exporter, nil
}
