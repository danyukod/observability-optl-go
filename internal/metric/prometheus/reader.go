package prometheus

import (
	"fmt"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

// Exporter is a Prometheus exporter.
func Reader() (metric.Reader, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus exporter: %w", err)
	}
	return exporter, nil
}
