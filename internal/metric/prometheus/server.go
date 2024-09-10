package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func ServeMetrics() {
	log.Printf("serving metrics at localhost:8181/metrics")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8181", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}
