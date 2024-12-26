package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, duration time.Duration) {
	defer wg.Done() // Ensure wg.Done() is called even if an error occurs
	fmt.Printf("Worker %d started\n", id)

	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d cancelled\n", id)
		return
	case <-time.After(duration):
		fmt.Printf("Worker %d finished\n", id)
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel the context when the main function exits

	// Start 3 workers with different durations
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter
		go worker(ctx, &wg, i, time.Duration(i)*time.Second)
	}

	// Simulate some work
	fmt.Println("Main thread doing some work...")
	time.Sleep(2 * time.Second)

	// Cancel the context, signaling workers to stop
	cancel()

	// Wait for all workers to complete
	fmt.Println("Main thread waiting for workers to finish...")
	wg.Wait()

	fmt.Println("All workers have finished. Main thread continues.")
}