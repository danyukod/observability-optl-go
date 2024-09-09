package pkg

import (
	"context"
	logger "github.com/danyukod/observability-optl-go/internal/log"
	"github.com/danyukod/observability-optl-go/internal/metric"
	"github.com/danyukod/observability-optl-go/internal/trace"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"log"
)

func InitObservability(ctx context.Context, name string) {
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
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %s", err)
		}
	}()

	shutdownMeterProvider, err := metric.InitMeterProvider(ctx, res)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownMeterProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown MeterProvider: %s", err)
		}
	}()

	shutdownLoggerProvider, err := logger.InitLoggerProvider(ctx, res)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownLoggerProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown LoggerProvider: %s", err)
		}
	}()
}
