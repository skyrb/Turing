package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, duration time.Duration) {
	defer wg.Done() // Ensure wg.Done() is called even if an error occurs
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(duration)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Start 3 workers with different durations
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter
		go worker(&wg, i, time.Duration(i)*time.Second)
	}

	// Wait for all workers to complete
	fmt.Println("Main thread waiting for workers to finish...")
	wg.Wait()

	fmt.Println("All workers have finished. Main thread continues.")
}