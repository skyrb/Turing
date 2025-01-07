package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	numWorkers := 3
	numJobs := 5
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	// Close the channel to signal workers to exit
	close(jobs)

	// Wait for workers to complete
	wg.Wait()
	fmt.Println("All jobs done.")

	// Simulate additional workload or program continuation
	time.Sleep(5 * time.Second)
	fmt.Println("Program ending.")
}