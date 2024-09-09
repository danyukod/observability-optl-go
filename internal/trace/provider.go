package trace

import (
	"context"
	"fmt"
	"github.com/danyukod/observability-optl-go/internal/trace/jaeger"
	"github.com/danyukod/observability-optl-go/internal/trace/otlp"
	"github.com/danyukod/observability-optl-go/internal/trace/stdout"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

func InitTraceProvider(ctx context.Context, res *resource.Resource) (func(context.Context) error, error) {
	traceExporter, err := initTraceExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tracerProvider.Shutdown, nil
}

func initTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	if os.Getenv("ENV") == "local" {
		return stdout.SpanExporter()
	}

	if os.Getenv("JAEGER") == "true" {
		return jaeger.SpanExporter(ctx)
	}

	return otlp.SpanExporter(ctx)
}
