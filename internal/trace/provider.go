package trace

import (
	"context"
	"fmt"
	"github.com/danyukod/observability-optl-go/internal/trace/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

func initTracerProvider(ctx context.Context, res *resource.Resource) (func(context.Context) error, error) {
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
	if os.Getenv("JAEGER") == "true" {
		return jaeger.SpanExporter(ctx)
	}

	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create otlp exporter: %w", err)
	}

	return exporter, nil
}
