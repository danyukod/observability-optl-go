package metric

import (
	"context"
	"fmt"
	"github.com/danyukod/observability-optl-go/internal/metric/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"os"
)

func initMeterProvider(ctx context.Context, res *resource.Resource) (func(context.Context) error, error) {
	metricExporter, err := initMetricReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(metricExporter),
		sdkmetric.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}

func initMetricReader(ctx context.Context) (sdkmetric.Reader, error) {
	if os.Getenv("PROMETHEUS") == "true" {
		return prometheus.PrometheusReader()
	}
	exporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}
	return sdkmetric.NewPeriodicReader(exporter), nil
}
