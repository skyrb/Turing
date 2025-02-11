package main

import (
	"fmt"
	"net/http"
	"time"
)

// Shared counter
var counter int

// Channel for increment requests
var incrementRequests chan bool

func main() {
	// Initialize the channel
	incrementRequests = make(chan bool, 100) // Buffered channel for performance

	// Start a goroutine to handle increment requests
	go incrementHandler()

	// Start the HTTP server
	http.HandleFunc("/increment", incrementCounter)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// Goroutine that handles increment requests
func incrementHandler() {
	for range incrementRequests {
		counter++ // Increment the counter safely
		fmt.Println("Counter incremented:", counter)
		
		// Simulate processing time
		time.Sleep(5 * time.Second)
	}
}

// Handler function to handle HTTP requests
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	incrementRequests <- true // Send a request to increment
	fmt.Fprintf(w, "Request received to increment the counter.\n")
}
