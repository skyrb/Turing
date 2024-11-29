package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    "syscall"
    "time"
)

const (
    numWorkers = 4
    port       = "8080"
)

var (
    wg     sync.WaitGroup
    workerChan = make(chan bool, numWorkers)
    shutdown   = make(chan struct{})
)

func main() {
    fmt.Printf("Starting server with %d workers on port %s\n", numWorkers, port)

    // Start worker goroutines
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker()
    }

    // Start the HTTP server
    go func() {
        if err := http.ListenAndServe(":"+port, nil); err != nil {
            log.Fatalf("Error starting HTTP server: %v", err)
        }
    }()

    // Graceful shutdown on SIGINT or SIGTERM
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
    fmt.Println("Received shutdown signal, shutting down...")
    close(shutdown)

    // Wait for all workers to finish
    wg.Wait()
    fmt.Println("Server shut down gracefully.")
}

func worker() {
    defer wg.Done()
    for {
        select {
        case <-shutdown:
            fmt.Println("Worker shutting down.")
            workerChan <- true
            return
        default:
            // Handle incoming requests here
            // For simplicity, we'll just simulate work
            time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
            fmt.Println("Worker processed a request.")
        }
    }
}
