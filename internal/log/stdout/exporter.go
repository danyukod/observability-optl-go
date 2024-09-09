package stdout

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	logsdk "go.opentelemetry.io/otel/sdk/log"
)

func Exporter(ctx context.Context) (logsdk.Exporter, error) {
	exporter, err := stdoutlog.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create stfoutlog exporter: %w", err)
	}
	return exporter, nil
}
