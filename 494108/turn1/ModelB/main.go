package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func main() {
	// Create a WaitGroup and set the number of goroutines to wait for
	var wg sync.WaitGroup
	numGoroutines := 5
	wg.Add(numGoroutines)

	// Start the worker goroutines
	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}

// Worker function
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work by sleeping for a random duration
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Printf("Goroutine %d sleeping for %v\n", id, sleepDuration)
	time.Sleep(sleepDuration)
}