package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Simulate a large dataset
	datasetSize := 1000000
	data := make([]int, datasetSize)

	// Populate the dataset with random numbers
	rand.Seed(42)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	// Initial capacity for the validData slice
	initialCapacity := 100000
	validData := make([]int, 0, initialCapacity)

	// Use a sync.Mutex to safely append to validData from multiple goroutines
	var mu sync.Mutex

	// Buffered error channel to prevent blocking
	errors := make(chan error, 100)

	// Validate data concurrently using goroutines
	var wg sync.WaitGroup

	for _, value := range data {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			if !isValid(v) {
				errors <- fmt.Errorf("Invalid data entry: %d", v)
				retry(v, errors)
			} else {
				// Acquire mutex to safely append
				mu.Lock()
				validData = append(validData, v)
				mu.Unlock()
			}
		}(value)
	}

	// Goroutine to close the errors channel after all validations are complete
	go func() {
		wg.Wait()
		close(errors)
	}()

	// Output the number of valid entries
	go func() {
		for err := range errors {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	time.Sleep(5 * time.Second) // Allow error processing to complete
	fmt.Printf("Number of valid entries: %d\n", len(validData))
}

// Simulate a data validation function
func isValid(value int) bool {
	// Replace this with actual validation logic
	return value%2 == 0
}

// Retry function with exponential backoff
func retry(value int, errors chan error) {
	maxRetries := 3
	backoff := 1 * time.Second
	for retryCount := 0; retryCount < maxRetries; retryCount++ {
		if isValid(value) {
			return
		}
		time.Sleep(backoff)
		backoff *= 2
	}
	errors <- fmt.Errorf("Failed to validate after %d retries: %d", maxRetries, value)
}
