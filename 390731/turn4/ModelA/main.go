package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// Simulate heavier work by increasing sleep time
	startTime := time.Now()
	time.Sleep(5 * time.Second) // Simulate work for 5 seconds
	endTime := time.Now()
	fmt.Printf("Worker %d finished, worked for %s\n", id, endTime.Sub(startTime))
}

func main() {
	numWorkers := 5

	// Demonstration of concurrency with limited CPU cores
	fmt.Println("\n=== Running with limited CPU cores (1 core) ===")
	startTime := time.Now()
	runtime.GOMAXPROCS(1) // Limit to 1 core
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time with 1 core: %v\n", executionTime)
	fmt.Println("All workers have finished with limited CPU cores")

	// Demonstration of concurrency and parallelism with all CPU cores
	fmt.Println("\n=== Running with all available CPU cores ===")
	startTime = time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available cores
	wg = sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	endTime = time.Now()
	executionTime = endTime.Sub(startTime)
	fmt.Printf("Execution time with %d cores: %v\n", runtime.NumCPU(), executionTime)
	fmt.Println("All workers have finished with all available CPU cores")
}