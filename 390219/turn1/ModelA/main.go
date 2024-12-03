package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Shared counter variable
var counter int

// Mutex to lock access to the counter
var counterMutex sync.Mutex

// Handler function to increment the counter
func incrementCounterHandler(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex before incrementing
	counterMutex.Lock()
	defer counterMutex.Unlock() // Unlock as soon as possible

	// Increment the counter
	counter++

	// Send response
	fmt.Fprintf(w, "Counter value: %d\n", counter)
}

func main() {
	// Start a HTTP server
	http.HandleFunc("/increment", incrementCounterHandler)

	// Start the server on a specified port
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}