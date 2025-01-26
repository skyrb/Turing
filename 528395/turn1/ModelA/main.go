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
    Err       error
    Timestamp time.Time
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

// Goroutine function
func fetchData(ctx context.Context, wg *sync.WaitGroup, dataCh chan<- StockData) {
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
                log.Println("Error fetching data:", err1, err2, err3)
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

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var wg sync.WaitGroup
    dataCh := make(chan StockData)

    wg.Add(1)
    go fetchData(ctx, &wg, dataCh)

    // Main processing loop
    go func() {
        for data := range dataCh {
            if data.Err != nil {
                log.Printf("Error received: %v\n", data.Err)
                continue
            }
            fmt.Printf("Received data at %v: Prices: %v, News: %v, Volumes: %v\n", data.Timestamp, data.Prices, data.News, data.Volumes)
        }
    }()

    // Run for a specific duration
    time.Sleep(10 * time.Second)
    cancel()

    wg.Wait()
    close(dataCh)
}