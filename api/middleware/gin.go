package middleware

import (
	"github.com/danyukod/observability-optl-go/internal/metric"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/metric"
	"time"
)

func MetricsMiddleware(serviceName string) gin.HandlerFunc {
	requestCounter, requestDuration, err := metric.InitMetricInstrumentation(serviceName)
	if err != nil {
		panic(err)
	}
	app := attribute.String("app", serviceName)

	return func(c *gin.Context) {

		startTime := time.Now()

		c.Next()

		path := attribute.String("path", c.Request.URL.Path)

		attributes := sdkmetric.WithAttributes(app, path)

		duration := time.Since(startTime).Seconds()
		requestCounter.Add(c.Request.Context(), 1, attributes)
		requestDuration.Record(c.Request.Context(), duration, attributes)
	}
}
