package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Function to simulate processing data
func processData(ctx context.Context, dataChannel <-chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case value, ok := <-dataChannel:
			if !ok {
				return
			}
			time.Sleep(time.Duration(value) * time.Millisecond)
			fmt.Printf("Worker %d processed data: %d\n", workerID, value)
		case <-ctx.Done():
			fmt.Printf("Worker %d: context canceled, terminating processing.\n", workerID)
			return
		}
	}
}

// Function to manage the number of workers based on load
func manageWorkers(dataChannel <-chan int, wg *sync.WaitGroup, maxWorkers, baseWorkers int) {
	// Initialize worker count and context
	workerCount := uint32(baseWorkers)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate data arriving at a varying rate
	go func() {
		for i := 1; true; i++ {
			select {
			case <-ctx.Done():
				return
			case dataChannel <- i:
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}
	}()

	// Loop to adjust worker count based on data arrival rate
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		// Determining new number of workers based on simulated load
		newWorkers := uint32(maxWorkers)
		// Replace with actual load monitoring logic here

		// Dynamically add or remove workers
		for newWorkers != workerCount {
			if newWorkers > workerCount {
				// Add a new worker
				workerID := atomic.AddUint32(&workerCount, 1)
				wg.Add(1)
				go processData(ctx, dataChannel, wg, int(workerID))
			} else {
				// Remove an existing worker
				if atomic.AddUint32(&workerCount, -1) == 0 {
					fmt.Println("All workers removed. Exiting.")
					return
				}
				ctx, cancel = context.WithCancel(context.Background())
			}
		}
	}
}

func main() {
	const maxWorkers = 10
	const baseWorkers = 5
	var wg sync.WaitGroup

	dataChannel := make(chan int)

	// Manage workers dynamically
	go manageWorkers(dataChannel, &wg, maxWorkers, baseWorkers)

	// Simulate waiting until the system finishes processing data
	wg.Wait()

	fmt.Println("All data streams processed.")
}