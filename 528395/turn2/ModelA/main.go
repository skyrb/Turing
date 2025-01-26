package main

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"
)

// Define a struct for the data
type StockData struct {
    Prices    interface{}
    News      interface{}
    Volumes   interface{}
    Timestamp time.Time
}

// Health status struct
type DataSourceStatus struct {
    Name   string
    IsHealthy bool
    LastError error
}

// Simulated API functions
func fetchRealTimePrices(ctx context.Context) (interface{}, error) {
    // Simulate API call
    time.Sleep(1 * time.Second)
    return "RealTimePrices", nil
}

func fetchMarketNews(ctx context.Context) (interface{}, error) {
    // Simulate API call
    time.Sleep(1 * time.Second)
    return "MarketNews", nil
}

func fetchTradingVolumes(ctx context.Context) (interface{}, error) {
    // Simulate API call
    time.Sleep(1 * time.Second)
    return "TradingVolumes", nil
}

// Goroutine function with health checks
func fetchData(ctx context.Context, wg *sync.WaitGroup, dataCh chan<- StockData, healthCh chan<- DataSourceStatus) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            return
        default:
            prices, err1 := fetchRealTimePrices(ctx)
            news, err2 := fetchMarketNews(ctx)
            volumes, err3 := fetchTradingVolumes(ctx)

            if err1 != nil || err2 != nil || err3 != nil {
                updateHealthStatus(healthCh, "Prices", err1)
                updateHealthStatus(healthCh, "News", err2)
                updateHealthStatus(healthCh, "Volumes", err3)
                continue
            }

            dataCh <- StockData{
                Prices:    prices,
                News:      news,
                Volumes:   volumes,
                Timestamp: time.Now(),
            }

            time.Sleep(1 * time.Second) // Simulate wait time between fetches
        }
    }
}

// Function to update health status
func updateHealthStatus(healthCh chan<- DataSourceStatus, sourceName string, err error) {
	if err != nil {
		healthCh <- DataSourceStatus{Name: sourceName, IsHealthy: false, LastError: err}
	} else {
		healthCh <- DataSourceStatus{Name: sourceName, IsHealthy: true, LastError: nil}
	}
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var wg sync.WaitGroup
    dataCh := make(chan StockData)
    healthCh := make(chan DataSourceStatus)

    wg.Add(1)
    go fetchData(ctx, &wg, dataCh, healthCh)

    // Goroutine to monitor health checks
    go func() {
        for status := range healthCh {
            if !status.IsHealthy {
                log.Printf("Data source %s is unhealthy. Error: %v\n", status.Name, status.LastError)
            } else {
                log.Printf("Data source %s is healthy.\n", status.Name)
            }
        }
    }()

    // Main processing loop
    go func() {
        for data := range dataCh {
            fmt.Printf("Received data at %v: Prices: %v, News: %v, Volumes: %v\n", data.Timestamp, data.Prices, data.News, data.Volumes)
        }
    }()

    // Run for a specific duration
    time.Sleep(10 * time.Second)
    cancel()

    wg.Wait()
    close(dataCh)
    close(healthCh)
}