package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, workChan chan int, resultsChan chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancellation signal, exiting...\n", id)
			return
		case work, ok := <-workChan:
			if !ok {
				// workChan closed, no more work to do
				return
			}
			time.Sleep(time.Duration(work) * time.Millisecond)
			resultsChan <- work
		}
	}
}

func main() {
	numWorkers := 5
	var wg sync.WaitGroup

	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())

	// Make sure to cancel the context when the main function returns
	defer cancel()

	// Channels for work and results
	workChan := make(chan int, 100)
	resultsChan := make(chan int, 100)

	// Start worker goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, &wg, i, workChan, resultsChan)
	}

	// Simulate generating work and sending it to the workers through workChan
	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(100 * time.Millisecond)
			workChan <- i
		}
		close(workChan) // Signal that no more work will be added
	}()

	// Simulate consuming results from the resultsChan
	go func() {
		for result := range resultsChan {
			fmt.Printf("Received result: %d\n", result)
		}
	}()

	// Wait for all workers to finish
	fmt.Println("Main thread waiting for workers to finish...")
	wg.Wait()

	fmt.Println("All workers have finished. Main thread continues.")
}
