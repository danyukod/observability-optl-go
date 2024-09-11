package metric

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func InitMetricInstrumentation(serviceName string) (metric.Int64Counter, metric.Float64Histogram, error) {
	meter := otel.Meter(serviceName)

	const totalHttpRequests = "http_request_total"
	const secondsHttpRequestDuration = "http_request_duration_seconds"

	counter, err := meter.Int64Counter(
		totalHttpRequests,
		metric.WithDescription("Total number of HTTP requests"),
	)
	if err != nil {
		return nil, nil, err
	}

	var buckets = []float64{0.1, 0.3, 1.5, 10.5}
	histogram, err := meter.Float64Histogram(
		secondsHttpRequestDuration,
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithExplicitBucketBoundaries(buckets...),
	)
	if err != nil {
		return nil, nil, err
	}

	return counter, histogram, nil
}
