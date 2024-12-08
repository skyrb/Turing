package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	goroutinesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_active_goroutines",
		Help: "Current number of active goroutines",
	})
)

func init() {
	prometheus.MustRegister(goroutinesGauge)
}

func main() {
	go monitorGoroutines()
	go simulateAlert()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :2112")
	http.ListenAndServe(":2112", nil)
}

func monitorGoroutines() {
	for {
		goroutinesGauge.Set(float64(runtime.NumGoroutine()))
		time.Sleep(15 * time.Second)
	}
}

func simulateAlert() {
	threshold := 20.0
	for {
		// Retrieve the current value of the goroutines gauge
		currentValue, _ := goroutinesGauge.Value()
		
		// Check if the value exceeds the threshold
		if currentValue > threshold {
			fmt.Printf("Alert: Goroutine count exceeded threshold (%d goroutines, threshold: %f)\n", runtime.NumGoroutine(), threshold)
		}

		// Check every minute
		time.Sleep(1 * time.Minute)
	}
}