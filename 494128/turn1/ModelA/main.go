package main

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPResponseLoggerMiddleware logs the HTTP request response and returns the status.
func HTTPResponseLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// call the next handler function
		status, err := next(w, r)
		if err != nil {
			fmt.Printf("Error handling request: %v\n", err)
		}

		// log the request details
		url := r.URL.String()
		method := r.Method
		timeTaken := time.Since(start)
		fmt.Printf("%s %s (%.3fs)\n", method, url, timeTaken.Seconds())

		// Let the middleware chain continue with the status
		// Indicating successful pass by using a named return value
		_, _ = w.Write([]byte("Middleware Response\n"))
	}
}

// SimpleHandler handles the actual HTTP request.
func SimpleHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	w.Write([]byte("Hello from SimpleHandler!\n"))
	return http.StatusOK, nil
}

func main() {
	// Use the middleware with a simple handler
	http.HandleFunc("/", HTTPResponseLoggerMiddleware(SimpleHandler))

	// Start the server
	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}