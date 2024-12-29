package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"math/rand"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	// Simulate concurrent processing of a request
	numTasks := 5
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Fprintf(w, "Task %d completed\n", taskID)
		}(i)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Fprintf(w, "All tasks completed\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
	}
}