package observability

import (
	"context"
	logger "github.com/danyukod/observability-optl-go/internal/log"
	"github.com/danyukod/observability-optl-go/internal/metric"
	"github.com/danyukod/observability-optl-go/internal/metric/prometheus"
	"github.com/danyukod/observability-optl-go/internal/trace"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"log"
)

func Init(ctx context.Context, name string) func() {
	serviceName := semconv.ServiceNameKey.String(name)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := trace.InitTraceProvider(ctx, res)
	if err != nil {
		log.Fatal(err)
	}

	shutdownMeterProvider, err := metric.InitMeterProvider(ctx, res)
	if err != nil {
		log.Fatal(err)
	}

	shutdownLoggerProvider, err := logger.InitLoggerProvider(ctx, res)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		if errS := shutdownTracerProvider; errS != nil {
			log.Fatalf("failed to shutdown TracerProvider: %s", errS)
		}
		if errS := shutdownMeterProvider; errS != nil {
			log.Fatalf("failed to shutdown MeterProvider: %s", errS)
		}
		if errS := shutdownLoggerProvider; errS != nil {
			log.Fatalf("failed to shutdown LoggerProvider: %s", errS)
		}
	}
}

func ServeMetrics() {
	channelError := make(chan error, 1)
	go prometheus.InitPrometheusServer(channelError)
	if err := <-channelError; err != nil {
		log.Fatalf("failed to start prometheus server: %s", err)
	}
}
