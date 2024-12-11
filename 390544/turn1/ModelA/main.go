package main

import (
	"context"
	"fmt"
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

// handler is the HTTP handler that uses the external service.
func handler(w http.ResponseWriter, r *http.Request) {
	// Set a timeout for our external call
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel() // Cancel the context before returning to avoid leaks

	// Call the external service
	result, err := simulateExternalCall(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			http.Error(w, "Request timed out. Please try again.", http.StatusGatewayTimeout)
		} else {
			http.Error(w, "An error occurred. Please try again later.", http.StatusInternalServerError)
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