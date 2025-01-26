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

// DataSourceConfig represents configuration for each data source
type DataSourceConfig struct {
    Name           string
    Priority       int
    Timeout        time.Duration
    RetryLimit     int
    RetryDelay     time.Duration
}

// Health status struct
type DataSourceStatus struct {
    Name      string
    IsHealthy bool
    LastError error
}

// Goroutine function with retry and timeout
func fetchData(ctx context.Context, cfg DataSourceConfig, dataCh chan<- StockData, healthCh chan<- DataSourceStatus, wg *sync.WaitGroup, supervisorCh chan<- string) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            return
        default:
            retries := 0
            for {
                if retries >= cfg.RetryLimit {
                    supervisorCh <- cfg.Name
                    break
                }
                cancellableCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
                data, err := simulateAPICall(cancellableCtx, cfg.Name)
                cancel()

                if err != nil {
                    updateHealthStatus(healthCh, cfg.Name, err)
                    retries++
                    time.Sleep(cfg.RetryDelay)
                    continue
                }

                dataCh <- StockData{
                    Prices:    data,
                    Timestamp: time.Now(),
                }
                updateHealthStatus(healthCh, cfg.Name, nil)
                time.Sleep(1 * time.Second) // Simulate wait time between fetches
                break
            }
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
    highPriorityDataCh := make(chan StockData)
    lowPriorityDataCh := make(chan StockData)
    healthCh := make(chan DataSourceStatus)
    supervisorCh := make(chan string)

    dataSources := []DataSourceConfig{
        {Name: "StockPrices", Priority: 1, Timeout: 2 * time.Second, RetryLimit: 3, RetryDelay: 500 * time.Millisecond},
        {Name: "MarketNews", Priority: 2, Timeout: 2 * time.Second, RetryLimit: 3, RetryDelay: 500 * time.Millisecond},
        // Add more data sources as needed
    }

    for _, ds := range dataSources {
        wg.Add(1)
        if ds.Priority == 1 {
            go fetchData(ctx, ds, highPriorityDataCh, healthCh, &wg, supervisorCh)
        } else {
            go fetchData(ctx, ds, lowPriorityDataCh, healthCh, &wg, supervisorCh)
        }
    }

    // Goroutine to monitor health checks and restart unhealthy data sources
    go func() {
        for dsName := range supervisorCh {
            log.Printf("Restarting data source %s\n", dsName)
            for _, ds := range dataSources {
                if ds.Name == dsName {
                    wg.Add(1)
                    go fetchData(ctx, ds, dataCh, healthCh, &wg, supervisorCh)
                    break