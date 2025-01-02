package main

import (
	"fmt"
	"sync"
	"strings"
)

// func uppercaseStringsConcurrently performs string manipulations concurrently.
func uppercaseStringsConcurrently(strings []string) ([]string, error) {
	if strings == nil {
		return nil, fmt.Errorf("input strings cannot be nil")
	}

	var wg sync.WaitGroup
	var results []string

	// Number of goroutines equals the number of strings to process
	wg.Add(len(strings))

	for _, str := range strings {
		go func(s string) {
			// Perform some manipulation, for example, convert to uppercase
			result := strings.ToUpper(s)

			// Lock before adding to the results slice
			defer wg.Done()
			results = append(results, result)
		}(str)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return results, nil
}

func main() {
	inputStrings := []string{"hello", "world", "go", "programming"}

	results, err := uppercaseStringsConcurrently(inputStrings)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Uppercased strings:")
	for _, result := range results {
		fmt.Println(result)
	}
}