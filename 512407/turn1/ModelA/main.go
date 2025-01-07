package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool defines a pool of worker goroutines
type WorkerPool struct {
	tasks    chan int // Channel for tasks
	workers  int     // Number of worker goroutines
	wg       sync.WaitGroup
}

// NewWorkerPool initializes a new worker pool
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		tasks:    make(chan int, workers*10), // Buffered channel to prevent blocking
		workers:  workers,
		wg:       sync.WaitGroup{},
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker is a goroutine that processes tasks
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.tasks {
		fmt.Printf("Worker %d processing task %d\n", wp.workers-wp.wg.Done(), task)
		time.Sleep(time.Duration(task) * time.Millisecond)
	}
}

// Submit adds a task to the worker pool
func (wp *WorkerPool) Submit(task int) {
	wp.tasks <- task
}

// Shutdown stops the worker pool
func (wp *WorkerPool) Shutdown() {
	close(wp.tasks)
	wp.wg.Wait()
}

func main() {
	wp := NewWorkerPool(3)
	wp.Start()

	// Submit tasks
	for i := 1; i <= 10; i++ {
		wp.Submit(i)
	}

	// Wait for all tasks to be completed
	wp.Shutdown()

	fmt.Println("All tasks are completed.")
}