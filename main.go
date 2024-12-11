package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a message type for sending updates to the state manager
type Update struct {
	Value int
}

// State Manager: This goroutine will manage the counter state
func stateManager(updates <-chan Update, wg *sync.WaitGroup) {
	defer wg.Done()

	counter := 0
	for update := range updates {
		// Update the state
		counter += update.Value
		fmt.Printf("Counter updated to: %d\n", counter)
	}
	fmt.Println("State manager: Finished processing updates.")
}

// Worker: This goroutine will perform operations that affect the state
func worker(id int, updates chan<- Update, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some work
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100) // Simulate work
		updates <- Update{Value: 1}         // Send an update to the state manager
		fmt.Printf("Worker %d incremented counter\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	updates := make(chan Update)

	// Start the state manager
	wg.Add(1)
	go stateManager(updates, &wg)

	// Start multiple workers
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, updates, &wg)
	}

	// Wait for workers to finish
	wg.Wait()

	// Close the updates channel after all workers are done
	close(updates)

	// Wait for state manager to finish processing updates
	wg.Wait()
	fmt.Println("All goroutines finished.")
}