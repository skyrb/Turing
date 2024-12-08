package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	goroutineCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "goroutine_count",
		Help: "Number of running goroutines.",
	})
)

func init() {
	// Register the goroutine counter with the Prometheus registry.
	prometheus.MustRegister(goroutineCount)

	// Start a Goroutine for metric updates
	go updateGoroutineCount(context.Background(), 10*time.Second)
}

func updateGoroutineCount(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get the current number of goroutines
			numGoroutines := runtime.NumGoroutine()
			log.Printf("Updated goroutine count: %d\n", numGoroutines)

			// Update the Prometheus counter
			goroutineCount.Set(float64(numGoroutines))

		case <-ctx.Done():
			log.Println("Goroutine count updater stopped.")
			return
		}
	}
}

func main() {
	// Start a goroutine to simulate work
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()

	// Register the Prometheus HTTP handler
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}