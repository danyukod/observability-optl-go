package log

import (
	"context"
	"fmt"
	"github.com/danyukod/observability-optl-go/internal/constants"
	"github.com/danyukod/observability-optl-go/internal/log/otlp"
	"github.com/danyukod/observability-optl-go/internal/log/stdout"
	"go.opentelemetry.io/otel/log/global"
	logsdk "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"os"
)

func InitLoggerProvider(ctx context.Context, res *resource.Resource) (func(context.Context) error, error) {

	exporter, err := initLoggerExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}

	processor := logsdk.NewBatchProcessor(exporter)

	logProvider := logsdk.NewLoggerProvider(
		logsdk.WithProcessor(processor),
		logsdk.WithResource(res),
	)

	global.SetLoggerProvider(logProvider)

	return logProvider.Shutdown, nil
}

func initLoggerExporter(ctx context.Context) (logsdk.Exporter, error) {
	if os.Getenv(constants.LoggerExporter) == constants.OpenTelemetry {
		return otlp.Exporter(ctx)
	}
	return stdout.Exporter(ctx)
}
