package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Shared counter and mutex
var counter int
var mux sync.Mutex

// Handler function to increment the counter
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mux.Lock()              // Guard the critical section
	counter++               // Increment the counter safely
	fmt.Fprintf(w, "Counter: %d\n", counter) // Print the result
	mux.Unlock()            // Release the lock
}

func main() {
	http.HandleFunc("/increment", incrementCounter)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}