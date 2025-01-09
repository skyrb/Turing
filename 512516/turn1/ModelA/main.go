package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Define a user record structure
type UserRecord struct {
	ID    int
	Value float64
}

// A function that simulates a complex calculation for each user record
func complexCalculation(record UserRecord) float64 {
	// Simulate work by sleeping for a random duration
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// Perform the complex calculation
	return record.Value * 100 + float64(record.ID)
}

func main() {
	// Number of worker goroutines
	numWorkers := 4
	// A channel to send work to the workers
	workCh := make(chan UserRecord, numWorkers*10)
	// A channel to receive results from the workers
	resultCh := make(chan float64)
	// A mutex to protect the result list
	var wg sync.WaitGroup
	// A counter for total processed records
	processedRecords := 0

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range workCh {
				// Process the record
				result := complexCalculation(record)
				resultCh <- result
				// Increment the processed records counter
				atomic.AddInt32(&processedRecords, 1)
			}
		}()
	}

	// Generate a large dataset of user records
	records := make([]UserRecord, 10000)
	for i := 0; i < len(records); i++ {
		records[i] = UserRecord{ID: i + 1, Value: float64(rand.Intn(1000))}
		workCh <- records[i]
	}

	// Close the work channel to signal workers to finish
	close(workCh)

	// Wait for all workers to finish
	wg.Wait()

	// Gather and display results
	fmt.Println("Processed", processedRecords, "records.")
	for result := range resultCh {
		fmt.Printf("Result: %.2f\n", result)
	}

	// Close the result channel
	close(resultCh)
}