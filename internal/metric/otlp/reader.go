package otlp

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func Reader(ctx context.Context) (sdkmetric.Reader, error) {
	exporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}
	return sdkmetric.NewPeriodicReader(exporter), nil
}
