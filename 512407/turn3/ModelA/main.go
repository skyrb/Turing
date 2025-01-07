package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	numWorkers := 5
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, jobs)
	}

	// Send jobs
	for i := 0; i < 10; i++ {
		jobs <- i
	}

	// Close the channel after sending all jobs
	close(jobs)

	// Wait for workers to complete
	wg.Wait()
	fmt.Println("All jobs done.")

	// Simulate additional workload or program continuation
	time.Sleep(5 * time.Second)
	fmt.Println("Program ending.")
}