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
	memoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_memory_alloc",
		Help: "Current memory allocation in bytes",
	})
	cpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_cpu_usage",
		Help: "Current CPU usage percentage",
	})
)

func init() {
	prometheus.MustRegister(goroutinesGauge)
	prometheus.MustRegister(memoryGauge)
	prometheus.MustRegister(cpuGauge)
}

func main() {
	go monitorSystem()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :2112")
	http.ListenAndServe(":2112", nil)
}

func monitorSystem() {
	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		goroutinesGauge.Set(float64(runtime.NumGoroutine()))
		memoryGauge.Set(float64(memStats.Alloc))

		// Placeholder: Add logic to calculate CPU usage
		cpuUsage := calculateCPUUsage()
		cpuGauge.Set(cpuUsage)

		time.Sleep(15 * time.Second)
	}
}

// calculateCPUUsage would be implemented to provide actual CPU usage data
func calculateCPUUsage() float64 {
	// Implement logic to calculate current CPU usage
	return 0.0 // For demonstration purposes
}