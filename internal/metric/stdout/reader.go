package stdout

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func Reader() (sdkmetric.Reader, error) {
	exporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}
	return sdkmetric.NewPeriodicReader(exporter), nil
}
