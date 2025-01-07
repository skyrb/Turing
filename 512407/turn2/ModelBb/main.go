package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		// Simulate work
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished task %d\n", id, task)
	}
}

func main() {
	numWorkers := 3
	tasks := make(chan int, 10)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Generate tasks and send them to the channel
	for i := 1; i <= 10; i++ {
		tasks <- i
	}

	// Close the tasks channel to signal the workers that there are no more tasks
	close(tasks)

	// Wait for all workers to complete their tasks
	wg.Wait()
	fmt.Println("All tasks done.")
}  