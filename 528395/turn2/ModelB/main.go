
package main  
import (  
    "context"
    "fmt"
    "log"
    "sync"
    "time"
)

type HealthStatus int
const (
    // Healthy indicates that the goroutine is running normally.
    Healthy HealthStatus = iota
    // Unhealthy indicates that the goroutine has encountered an error.
    Unhealthy
)

// Define a struct for the data and goroutine state
type GoroutineState struct {
    Data      StockData
    Health    HealthStatus
    LastError error
}

func fetchData(ctx context.Context, wg *sync.WaitGroup, dataCh chan<- GoroutineState, errCh chan<- error, name string) {
    defer wg.Done()

    state := GoroutineState{Health: Healthy}

    for {
        select {
        case <-ctx.Done():
            return
        default:
            prices, err1 := fetchRealTimePrices(ctx)
            news, err2 := fetchMarketNews(ctx)
            volumes, err3 := fetchTradingVolumes(ctx)
            var err error

            if err1 != nil {
                err = err1
            }
            if err2 != nil {
                err = fmt.Errorf("%v: %w", err2, err)
            }
            if err3 != nil {
                err = fmt.Errorf("%v: %w", err3, err)
            }

            state.Data = StockData{
                Prices:    prices,
                News:      news,
                Volumes:   volumes,
                Err:       err,
                Timestamp: time.Now(),
            }

            if err != nil {
                state.Health = Unhealthy
                state.LastError = err
                log.Println("Error fetching data from", name, ":", err)
            } else {
                state.Health = Healthy
            }

            dataCh <- state

            time.Sleep(1 * time.Second) // Simulate wait time between fetches
        }
    }
}

func main() {  
    ctx, cancel := context.WithCancel(context.Background())  
    defer cancel()
    
    var wg sync.WaitGroup
    dataCh := make(chan GoroutineState)
    errCh := make(chan error)

    wg.Add(3)  
    go fetchData(ctx, &wg, dataCh, errCh, "Prices")
    go fetchData(ctx, &wg, dataCh, errCh, "News")
    go fetchData(ctx, &wg, dataCh, errCh, "Volumes")

    // Health check routine
    go func() {
        ticker := time.NewTicker(3 * time.Second) // Health check interval
        defer ticker.Stop()

        for {
            select {
            case <-ctx.Done():
                return
            case <-ticker.C:
                log.Println("Health Check:")
                for i := 0; i < 3; i++ {
                    state := <-dataCh
                    if state.Health == Unhealthy {
                        log.Printf("Goroutine %d is unhealthy: %v\n", i, state.LastError)
                    } else {
                        log.Println("Goroutine", i, "is healthy")
                    }
                }
            }
        }
    }()

    // Error handling routine
    go func() {
        for err := range errCh {
            log.Println("Received error:", err)
        }
    }()

    // Main processing loop
    go func() {
        for state := range dataCh {
            if state.Health == Healthy {
                fmt.Printf("Received data at %v: Prices: %v, News: %v, Volumes: %v\n", state.Data.Timestamp, state.Data.Prices, state.Data.News, state.Data.Volumes)
            }
        }