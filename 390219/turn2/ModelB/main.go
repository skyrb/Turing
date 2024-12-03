package main

import (
	"fmt"
	"net/http"
)

// Counter operation types
type CounterOp struct {
	Increment bool
	Response  chan int
}

func main() {
	counter := make(chan CounterOp)

	// Goroutine to manage counter operations
	go func() {
		var count int
		for op := range counter {
			if op.Increment {
				count++
			}
			op.Response <- count
		}
	}()

	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		response := make(chan int)
		counter <- CounterOp{Increment: true, Response: response}
		currentCount := <-response
		fmt.Fprintf(w, "Counter: %d\n", currentCount)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}