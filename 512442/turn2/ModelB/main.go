package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"sort"
)

type ErrorType int

const (
	CriticalError ErrorType = iota
	MinorError
)

type AggregatedError struct {
	Error    error
	ErrorType ErrorType
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// fetchData simulates fetching data from a URL and returns an AggregatedError.
func fetchData(url string) AggregatedError {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	// Simulate network request
	resp, err := http.Get(url)
	if err != nil {
		return AggregatedError{Error: err, ErrorType: CriticalError}
	}
	defer resp.Body.Close()

	// Simulate occasional error
	if rand.Intn(10) == 0 {
		return AggregatedError{Error: fmt.Errorf("simulated error fetching data from %s", url), ErrorType: MinorError}
	}

	return AggregatedError{Error: nil, ErrorType: CriticalError}
}

func main() {
	// URLs to fetch data from
	urls := []string{
		"http://example.com",
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://non-existent-url.com",
	}

	// Make a channel to send aggregated errors
	errorChan := make(chan AggregatedError, len(urls))

	// Synchronization to ensure we wait for all Goroutines to finish
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := fetchData(url)
			errorChan <- err
		}(url)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	// Close the error channel
	close(errorChan)

	// Collect and aggregate errors
	var errors []AggregatedError
	for err := range errorChan {
		errors = append(errors, err)
	}

	// Prioritize errors by type (critical first)
	sort.Slice(errors, func(i, j int) bool {
		return errors[i].ErrorType > errors[j].ErrorType
	})

	// Handle errors coherently
	for _, err := range errors {
		if err.ErrorType == CriticalError {
			fmt.Println("Critical Error:", err.Error)
			// Handle critical error here, e.g., retry, restart, etc.
		} else {
			fmt.Println("Minor Error:", err.Error)
			// Handle minor error here, e.g., log, ignore, etc.
		}
	}

	// Check if there were any errors
	if len(errors) > 0 {
		fmt.Println("Some errors occurred during data fetch.")
	} else {
		fmt.Println("All data fetched successfully.")
	}
}