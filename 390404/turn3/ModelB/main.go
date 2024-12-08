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
	alertThreshold = 2 // Set this to your desired threshold
)

func init() {
	prometheus.MustRegister(goroutinesGauge)
}

func main() {
	go monitorGoroutines()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :2112")
	http.ListenAndServe(":2112", nil)
}

func monitorGoroutines() {
	for {
		goroutines := runtime.NumGoroutine()
		goroutinesGauge.Set(float64(goroutines))
		checkAlert(goroutines)

		time.Sleep(5 * time.Second)
	}
}

func checkAlert(goroutines int) {
	if goroutines > alertThreshold {
		fmt.Printf("ALERT: Goroutine count exceeded threshold! Current: %d\n", goroutines)
		// Trigger local action like sending to a log file, webhook, email notification, etc.
	}
}