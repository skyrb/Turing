package main

import (
    "context"
    "fmt"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

const (
    maxWorkers = 10
    rateLimit  = 100 // Requests per second
)

func worker(ctx context.Context, wg *sync.WaitGroup, workChan <-chan int, resultsChan chan<- int, limiter *rate.Limiter) {
    defer wg.Done()
    for work := range workChan {
        if err := limiter.Wait(ctx); err != nil {
            fmt.Println("Worker: Rate limit exceeded, canceling task...")
            return
        }

        select {
        case <-ctx.Done():
            fmt.Println("Worker: Received cancellation signal, exiting...")
            return
        default:
            // Simulate work
            time.Sleep(time.Duration(work) * time.Millisecond)
            resultsChan <- work
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var wg sync.WaitGroup
    workChan := make(chan int, 100)
    resultsChan := make(chan int, 100)
    limiter := rate.NewLimiter(rate.Limit(rateLimit), 1)

    for i := 0; i < maxWorkers; i++ {
        wg.Add(1)
        go worker(ctx, &wg, workChan, resultsChan, limiter)
    }

    go func() {
        for i := 1; i <= 1000; i++ {
            time.Sleep(time.Second)
            workChan <- i
        }
        close(workChan)
    }()

    go func() {
        for result := range resultsChan {
            fmt.Printf("Received result: %d\n", result)
        }
    }()

    fmt.Println("Main thread waiting for workers to finish...")
    wg.Wait()
    fmt.Println("All workers have finished. Main thread continues.")

    // Graceful shutdown
    fmt.Println("Shutting down workers...")
    cancel()
}  