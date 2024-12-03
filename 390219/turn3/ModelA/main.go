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
	incrementRequests = make(chan bool)

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

		// Simulate additional processing
		fmt.Println("Processing started for counter value:", counter)
		time.Sleep(2 * time.Second) // Simulate processing time

		// Wait for 5 seconds before accepting the next increment request
		time.Sleep(5 * time.Second)
		fmt.Println("Waiting period completed for counter value:", counter)
	}
}

// Handler function to handle HTTP requests
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	incrementRequests <- true // Send a request to increment
	fmt.Fprintf(w, "Counter: %d\n", counter) // Send the current counter value
}