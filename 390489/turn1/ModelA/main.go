package main

import (
	"fmt"
	"net/http"
	"time"
)

// simulateAPICall simulates an API call that takes some time to complete.
func simulateAPICall(url string, duration time.Duration) (*http.Response, error) {
	fmt.Printf("Simulating call to %s for %v\n", url, duration)
	time.Sleep(duration)

	// For demonstration purposes, return a successful response
	resp := &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Body:       nil,
	}
	return resp, nil
}

func main() {
	urls := []string{"http://example.com/api/data1", "http://example.com/api/data2"}
	timeouts := []time.Duration{5 * time.Second, 2 * time.Second}

	for i, url := range urls {
		timeout := timeouts[i]
		go func(url string, timeout time.Duration) {
			// Use a context to manage the timeout
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			// Simulate the API call
			resp, err := simulateAPICall(url, 3*time.Second)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}
			defer resp.Body.Close()

			// Process the response
			fmt.Printf("Response from %s: StatusCode=%d\n", url, resp.StatusCode)
		}(url, timeout)
	}

	// Wait for the goroutines to complete
	time.Sleep(6 * time.Second)
}