package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	
	for i := 0; i < iterations; i++ {
		// Perform some computation
		_ = i * i * i
	}
	
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	numWorkers := 10
	iterations := 1_000_000_000 // 1 billion iterations for more intensive work

	// Demonstration of concurrency with limited CPU cores
	fmt.Println("\n=== Running with limited CPU cores (1 core) ===")
	startTime := time.Now()
	runtime.GOMAXPROCS(1) // Limit to 1 core
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, iterations)
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
		go worker(i, &wg, iterations)
	}
	wg.Wait()
	endTime = time.Now()
	executionTime = endTime.Sub(startTime)
	fmt.Printf("Execution time with %d cores: %v\n", runtime.NumCPU(), executionTime)
	fmt.Println("All workers have finished with all available CPU cores")
}