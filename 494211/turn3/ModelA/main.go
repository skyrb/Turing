package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// func uppercaseStringsConcurrentlyWithMetrics performs string manipulations concurrently, returning results and execution metrics.
func uppercaseStringsConcurrentlyWithMetrics(input []string) ([]string, []time.Duration) {
	if input == nil {
		return nil, nil
	}

	var wg sync.WaitGroup
	results := make([]string, len(input))
	durations := make([]time.Duration, len(input))
	var mu sync.Mutex

	// Number of goroutines equals the number of strings to process
	wg.Add(len(input))

	for i, str := range input {
		go func(i int, s string) {
			defer wg.Done()

			// Measure execution time
			start := time.Now()

			// Perform some manipulation, for example, convert to uppercase
			result := strings.ToUpper(s)

			// Calculate execution time
			duration := time.Since(start)

			// Lock before updating results and durations slices
			mu.Lock()
			results[i] = result
			durations[i] = duration
			mu.Unlock()
		}(i, str)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return results, durations
}

func main() {
	inputStrings := []string{"hello", "world", "go", "programming"}

	results, durations := uppercaseStringsConcurrentlyWithMetrics(inputStrings)

	fmt.Println("\nTransformed results:")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("\nExecution times:")
	for i, duration := range durations {
		fmt.Printf("String %d execution time: %.6f ns\n", i, duration.Nanoseconds())
	}
}