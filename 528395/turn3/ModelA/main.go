package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"math/rand"
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

// Simulated API function
func simulateAPICall(ctx context.Context, name string) (interface{}, error) {
	// Simulate varied response times and occasional errors
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(time.Duration(1+rand.Intn(3)) * time.Second):
	}
	// Simulate 10% failure rate
	if rand.Float32() < 0.1 {
		return nil, fmt.Errorf("%s failed to respond", name)
	}
	return fmt.Sprintf("%sData", name), nil
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
				prices, err := simulateAPICall(cancellableCtx, cfg.Name)
				cancel()

				if err != nil {
					updateHealthStatus(healthCh, cfg.Name, err)
					retries++
					time.Sleep(cfg.RetryDelay)
					continue
				}

				dataCh <- StockData{
					Prices:    prices,
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
	dataCh := make(chan StockData)
	healthCh := make(chan DataSourceStatus)
	supervisorCh := make(chan string)

	dataSources := []DataSourceConfig{
		{Name: "DataSource1", Timeout: 2 * time.Second, RetryLimit: 3, RetryDelay: 500 * time.Millisecond},
		{Name: "DataSource2", Timeout: 2 * time.Second, RetryLimit: 3, RetryDelay: 500 * time.Millisecond},
		// Add more data sources as needed
	}

	for _, ds := range dataSources {
		wg.Add(1)
		go fetchData(ctx, ds, dataCh, healthCh, &wg, supervisorCh)
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
				}
			}
		}
	}()

	// Goroutine to handle data and health status processing
	go func() {
		for {
			select {
			case data := <-dataCh:
				fmt.Printf("Received data at %v: Prices: %v\n", data.Timestamp, data.Prices)
			case status := <-healthCh:
				if !status.IsHealthy {
					log.Printf("Data source %s is unhealthy. Error: %v\n", status.Name, status.LastError)
				} else {
					log.Printf("Data source %s is healthy.\n", status.Name)
				}
			}
		}
	}()

	// Run for a specific duration
	time.Sleep(30 * time.Second)
	cancel()

	wg.Wait()
	close(dataCh)
	close(healthCh)
	close(supervisorCh)
}