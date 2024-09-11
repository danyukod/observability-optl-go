package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitPrometheusServer(channelError chan error) {
	router := gin.Default()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	err := router.Run(":8181")
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		channelError <- err
		return
	}
}
