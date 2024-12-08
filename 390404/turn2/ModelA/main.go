package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
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

var (
	prevGoroutines int = runtime.NumGoroutine()
	updateInterval   time.Duration = 5 * time.Second
	intervalMutex    sync.Mutex
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
		currentGoroutines := runtime.NumGoroutine()
		goroutinesGauge.Set(float64(currentGoroutines))

		// Calculate change in goroutine count
		change := abs(currentGoroutines - prevGoroutines)
		prevGoroutines = currentGoroutines

		// Adjust update interval based on change
		if change > 10 {
			intervalMutex.Lock()
			updateInterval = 1 * time.Second
			intervalMutex.Unlock()
		} else if change < 2 {
			intervalMutex.Lock()
			updateInterval = 10 * time.Second
			intervalMutex.Unlock()
		}

		// Sleep for the current update interval
		time.Sleep(updateInterval)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}