package metric

import (
	"context"
	"fmt"
	"github.com/danyukod/observability-optl-go/internal/constants"
	"github.com/danyukod/observability-optl-go/internal/metric/otlp"
	"github.com/danyukod/observability-optl-go/internal/metric/prometheus"
	"github.com/danyukod/observability-optl-go/internal/metric/stdout"
	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"os"
)

func InitMeterProvider(ctx context.Context, res *resource.Resource) (func(context.Context) error, error) {
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
	if os.Getenv(constants.MetricReader) == constants.Prometheus {
		return prometheus.Reader()
	}
	if os.Getenv(constants.MetricReader) == constants.OpenTelemetry {
		return otlp.Reader(ctx)
	}
	return stdout.Reader()
}
