package metric

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func InitMetricInstrumentation(serviceName string) (metric.Int64Counter, metric.Float64Histogram, error) {
	meter := otel.Meter(serviceName)

	counter, err := meter.Int64Counter("http_request_total", metric.WithDescription("Total number of HTTP requests"))
	if err != nil {
		return nil, nil, err
	}

	histogram, err := meter.Float64Histogram("http_request_duration_seconds", metric.WithDescription("HTTP request duration in seconds"))
	if err != nil {
		return nil, nil, err
	}

	return counter, histogram, nil
}
