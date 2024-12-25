package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processData(data <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range data {
		select {
		case <-ctx.Done():
			fmt.Println("Received cancellation signal. Terminating data processing.")
			return
		default:
			time.Sleep(time.Duration(value) * time.Millisecond)
			fmt.Println("Processed data:", value)
		}
	}
}

func main() {
	const numGoroutines = 5
	var wg sync.WaitGroup
	dataStream1 := make(chan int)
	dataStream2 := make(chan int)

	// Start worker Goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processData(dataStream1, context.TODO(), &wg)
		go processData(dataStream2, context.TODO(), &wg)
	}

	// Populate data streams
	go func() {
		for i := 1; i <= 10; i++ {
			dataStream1 <- i
		}
		close(dataStream1)
	}()

	go func() {
		for i := 11; i <= 20; i++ {
			dataStream2 <- i
		}
		close(dataStream2)
	}()

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wait for all processing to complete or for the timeout to occur
	done := make(chan bool, numGoroutines)
	go func() {
		<-ctx.Done()
		cancel() // Cancel any remaining operations if the timeout occurs
		close(done)
	}()

	select {
	case <-done: // Wait for context cancellation
		fmt.Println("Context timed out. Stopping data processing.")
	case <-wg.Done(): // Wait for all workers to finish
	}

	fmt.Println("All data streams processed or interrupted.")
}