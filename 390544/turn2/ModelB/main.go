package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// simulateExternalCall simulates an external API call.
func simulateExternalCall(ctx context.Context) (string, error) {
	// Simulate processing time
	select {
	case <-time.After(2 * time.Second): // Simulate a long operation
		return "Success", nil
	case <-ctx.Done(): // If the context was canceled or timed out
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

// handler is the HTTP handler that uses the external service with retries.
func handler(w http.ResponseWriter, r *http.Request) {
	// Set a timeout for our external call
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel() // Cancel the context before returning to avoid leaks

	// Define retry parameters
	maxAttempts := 3
	backoff := 100 * time.Millisecond

	// Call the external service with retries
	result, err := retry(ctx, maxAttempts, backoff, func(ctx context.Context) (string, error) {
		return simulateExternalCall(ctx)
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
		fmt.Println("Failed to start server:", err)
	}
}