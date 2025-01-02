package main

import (
	"fmt"
	"strings"
	"sync"
)

// func uppercaseStringsConcurrentlyWithErrors processes strings concurrently and collects results with errors.
func uppercaseStringsConcurrentlyWithErrors(input []string) ([]string, []error) {
	if input == nil {
		return nil, []error{fmt.Errorf("input strings cannot be nil")}
	}

	var wg sync.WaitGroup
	var results []string
	var errors []error

	// Number of goroutines equals the number of strings to process
	wg.Add(len(input))

	for _, str := range input {
		go func(s string) {
			defer wg.Done()

			// Simulate an error for the purpose of this example
			if strings.Contains(s, "error") {
				errors = append(errors, fmt.Errorf("failed to process '%s': error encountered", s))
				return
			}

			// Perform some manipulation, for example, convert to uppercase
			result := strings.ToUpper(s)

			// Append result and nil error to the respective slices
			results = append(results, result)
			errors = append(errors, nil)
		}(str)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Extract the successful results
	successfulResults := make([]string, 0, len(results))
	for i, err := range errors {
		if err == nil {
			successfulResults = append(successfulResults, results[i])
		} else {
			// Log error for problematic string
			fmt.Printf("Error: %v\n", err)
		}
	}

	return successfulResults, errors
}

func main() {
	inputStrings := []string{"hello", "world", "go", "error", "programming"}

	results, errors := uppercaseStringsConcurrentlyWithErrors(inputStrings)

	fmt.Println("Uppercased strings (successful only):")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("Errors encountered:")
	for _, err := range errors {
		if err != nil {
			fmt.Println(err)
		}
	}
}