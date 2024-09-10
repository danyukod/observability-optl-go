package middleware

import (
	"github.com/danyukod/observability-optl-go/internal/metric"
	"github.com/gin-gonic/gin"
	"time"
)

func MetricsMiddleware(serviceName string) gin.HandlerFunc {
	requestCounter, requestDuration, err := metric.InitMetricInstrumentation(serviceName)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime).Seconds()
		requestCounter.Add(c.Request.Context(), 1)
		requestDuration.Record(c.Request.Context(), duration)
	}
}
