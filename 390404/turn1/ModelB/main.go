package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Declare a gauge metric to track the number of goroutines
var (
	goroutinesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_active_goroutines",
		Help: "Current number of active goroutines",
	})
)

func init() {
	// Register the goroutines gauge with Prometheus
	prometheus.MustRegister(goroutinesGauge)
}

func main() {
	// Start a goroutine to update the goroutine count periodically
	go monitorGoroutines()

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :2112")
	http.ListenAndServe(":2112", nil)
}

func monitorGoroutines() {
	for {
		// Update the gauge with the current number of goroutines
		goroutinesGauge.Set(float64(runtime.NumGoroutine()))

		// Sleep for a specified duration before the next update
		time.Sleep(5 * time.Second)
	}
}