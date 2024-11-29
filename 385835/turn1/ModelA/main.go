package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/exp/rand"
)

const (
	numWorkers = 4
	port       = "8080"
)

var (
	wg         sync.WaitGroup
	workerChan = make(chan struct{}, numWorkers)
	shutdown   = make(chan struct{})
)

func main() {
	fmt.Printf("Starting server with %d workers on port %s\n", numWorkers, port)

	// Start worker Goroutines
	for i := 0; i < numWorkers; i++ {
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
	wg.Add(1) // Add one to wait group for each worker
	defer wg.Done()

	for {
		select {
		case <-shutdown:
			fmt.Println("Worker shutting down.")
			return
		default:
			// Handle incoming requests here
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Println("Worker processed a request.")
		}
	}
}
