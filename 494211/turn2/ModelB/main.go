package main

import (
	"fmt"
	"strings"
	"sync"
)

// func uppercaseStringsConcurrently performs string manipulations concurrently, returning both results and errors.
func uppercaseStringsConcurrently(input []string) ([]string, []error) {
	if input == nil {
		return nil, []error{fmt.Errorf("input strings cannot be nil")}
	}

	var wg sync.WaitGroup
	resultChan := make(chan struct {
		result string
		err    error
	})

	// Number of goroutines equals the number of strings to process
	wg.Add(len(input))

	for _, str := range input {
		go func(s string) {
			defer wg.Done()

			// Perform some manipulation, for example, convert to uppercase
			result, err := processString(s)

			// Send the result and error through the channel
			resultChan <- struct {
				result string
				err    error
			}{result, err}
		}(str)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []string
	var errors []error

	// Collect results and errors from the channel
	for result := range resultChan {
		if result.err != nil {
			errors = append(errors, result.err)
		} else {
			results = append(results, result.result)
		}
	}

	return results, errors
}

// processString performs the string manipulation (in this case, converting to uppercase) and returns an error if applicable.
func processString(s string) (string, error) {
	// Add some logic to generate an error for certain strings
	if s == "go" {
		return "", fmt.Errorf("processing string '%s' failed", s)
	}
	return strings.ToUpper(s), nil
}

func main() {
	inputStrings := []string{"hello", "world", "go", "programming"}

	results, errors := uppercaseStringsConcurrently(inputStrings)

	fmt.Println("Successful results:")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("\nEncountered errors:")
	for _, err := range errors {
		fmt.Println(err)
	}
}  