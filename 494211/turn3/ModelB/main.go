package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// resultAndMetrics holds the result and processing metrics for each goroutine.
type resultAndMetrics struct {
	Result string
	Time   time.Duration
	Err    error
}

func manipulateStringsConcurrently(input []string, numWorkers int) ([]string, map[string]time.Duration, []error) {
	if input == nil {
		return nil, nil, []error{fmt.Errorf("input strings cannot be nil")}
	}

	if numWorkers <= 0 {
		numWorkers = 1
	}

	var wg sync.WaitGroup
	jobChan := make(chan string)
	resultChan := make(chan resultAndMetrics)

	// Start worker goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for str := range jobChan {
				startTime := time.Now()
				result, err := processString(str)
				endTime := time.Now()
				duration := endTime.Sub(startTime)
				resultChan <- resultAndMetrics{Result: result, Time: duration, Err: err}
			}
		}()
	}

	// Dispatch jobs to worker goroutines
	go func() {
		for _, str := range input {
			jobChan <- str
		}
		close(jobChan)
	}()

	// Collect results and metrics from worker goroutines
	var results []string
	metrics := make(map[string]time.Duration)
	var errors []error
	go func() {
		for result := range resultChan {
			if result.Err != nil {
				errors = append(errors, result.Err)
			} else {
				results = append(results, result.Result)
				metrics[result.Result] = result.Time
			}
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
	close(resultChan)

	return results, metrics, errors
}

// processString performs the string manipulation (in this case, converting to uppercase) and returns an error if applicable.
func processString(s string) (string, error) {
	// Add some logic to generate an error for certain strings
	if s == "go" {
		return "", fmt.Errorf("processing string '%s' failed", s)
	}
	time.Sleep(time.Duration(100+len(s)*10) * time.Millisecond) // Add some delay for demonstration
	return strings.ToUpper(s), nil
}

func main() {
	inputStrings := []string{"hello", "world", "go", "programming", "is", "fun"}
	numWorkers := 3

	results, metrics, errors := manipulateStringsConcurrently(inputStrings, numWorkers)

	fmt.Println("Successful results:")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("\nProcessing metrics:")
	for result, duration := range metrics {
		fmt.Printf("String: %s, Time: %s\n", result, duration)
	}

	fmt.Println("\nEncountered errors:")
	for _, err := range errors {
		fmt.Println(err)
	}
}