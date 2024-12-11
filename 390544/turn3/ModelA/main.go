package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// simulateExternalCall simulates different durations for different API calls.
func simulateExternalCall(ctx context.Context, endpoint string) (string, error) {
	var duration time.Duration

	// Set duration based on the endpoint type
	switch endpoint {
	case "quick":
		duration = 500 * time.Millisecond // Simulate a quick call
	case "normal":
		duration = 2 * time.Second // Simulate a normal call
	case "slow":
		duration = 5 * time.Second // Simulate a slow call
	default:
		return "", fmt.Errorf("unknown endpoint")
	}

	// Simulate processing time
	select {
	case <-time.After(duration):
		return "Success", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// retry calls a function with a specified number of retries and a backoff strategy.
func retry(ctx context.Context, attempts int, backoff time.Duration, f func(context.Context) (string, error)) (string, error) {
	for attempt := 1; attempt <= attempts; attempt++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			result, err := f(ctx)
			if err == nil {
				return result, nil
			}
			log.Printf("Attempt %d failed: %v. Retrying in %v...", attempt, err, backoff)
			time.Sleep(backoff)
			backoff *= 2 // Exponential backoff
		}
	}
	return "", fmt.Errorf("max retries reached")
}

// handler is the HTTP handler that uses the external service with retries and dynamic timeout.
func handler(w http.ResponseWriter, r *http.Request) {
	// Determine timeout based on the endpoint type
	endpoint := r.URL.Query().Get("type") // Assume 'type' parameter indicates the API type
	var timeout time.Duration

	switch endpoint {
	case "quick":
		timeout = 1 * time.Second
	case "normal":
		timeout = 3 * time.Second
	case "slow":
		timeout = 6 * time.Second
	default:
		http.Error(w, "Invalid type parameter", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel() // Cancel the context before returning to avoid leaks

	// Define retry parameters
	maxAttempts := 3
	backoff := 100 * time.Millisecond

	// Call the external service with retries
	result, err := retry(ctx, maxAttempts, backoff, func(ctx context.Context) (string, error) {
		return simulateExternalCall(ctx, endpoint)
	})

	if err != nil {
		if err == context.DeadlineExceeded {
			http.Error(w, "Request timed out. Please try again.", http.StatusGatewayTimeout)
		} else {
			http.Error(w, fmt.Sprintf("An error occurred: %v. Please try again later.", err), http.StatusInternalServerError)
		}
		return
	}

	// If successful, return the result
	fmt.Fprintf(w, "Result: %s", result)
}

func main() {
	http.HandleFunc("/api", handler)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}