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

// retryWithTimeout retries the given function fn until it succeeds or the context fails.
func retryWithTimeout(ctx context.Context, fn func(context.Context) (string, error), maxAttempts int, delay time.Duration) (string, error) {
	var lastErr error
	for attempts := 0; attempts < maxAttempts; attempts++ {
		result, err := fn(ctx)
		if err == nil {
			return result, nil // Success
		}
		lastErr = err
		log.Printf("Attempt %d failed: %v", attempts+1, err)
		time.Sleep(delay) // Wait before retrying
	}
	return "", lastErr
}

// handler is the HTTP handler that uses the external service.
func handler(w http.ResponseWriter, r *http.Request) {
	// Set a timeout for our external call
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel() // Cancel the context before returning to avoid leaks

	// Retry calling the external service
	result, err := retryWithTimeout(ctx, simulateExternalCall, 3, 500*time.Millisecond)
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
		log.Println("Failed to start server:", err)
	}
}