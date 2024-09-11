package main

import (
	"context"
	"github.com/danyukod/observability-optl-go/api/middleware"
	"github.com/danyukod/observability-optl-go/pkg"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

const serviceName = "observability-optl-go"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdownObservability := observability.Init(ctx, serviceName)
	defer shutdownObservability()

	//Start the prometheus HTTP server and pass the exporter Collector to it
	go observability.ServeMetrics()

	// Set up Gin router
	router := gin.Default()

	router.Use(middleware.MetricsMiddleware(serviceName))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")

}
