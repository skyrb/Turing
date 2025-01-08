package main

import (
	"fmt"
	"sync"
)

// SalesData represents sales data from a single marketplace.
type SalesData struct {
	Marketplace string
	Revenue      float64
}

// WorkerPool manages a fixed number of worker goroutines.
type WorkerPool struct {
	tasks     chan *SalesData
	workers   int
	wg        *sync.WaitGroup
	shutdown  chan struct{}
	totalRevenue float64
	mu           sync.Mutex
}

// NewWorkerPool creates a new worker pool with the specified number of workers.
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		tasks:     make(chan *SalesData, 100), // Buffered channel
		workers:   workers,
		wg:        &sync.WaitGroup{},
		shutdown:  make(chan struct{}),
		totalRevenue: 0.0,
	}
}

// Start starts the worker pool.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker()
	}
}

// worker is a single worker goroutine.
func (wp *WorkerPool) worker() {
	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok {
				return // Worker pool is shutting down
			}
			wp.processTask(task)
		case <-wp.shutdown:
			return // Worker pool is shutting down
		}
	}
}

// processTask processes a sales data task and accumulates the total revenue.
func (wp *WorkerPool) processTask(task *SalesData) {
	fmt.Printf("Processing sales data from %s: $%.2f\n", task.Marketplace, task.Revenue)
	wp.mu.Lock()
	wp.totalRevenue += task.Revenue
	wp.mu.Unlock()
	wp.wg.Done()
}

// Submit submits a task to the worker pool.
func (wp *WorkerPool) Submit(task *SalesData) {
	wp.tasks <- task
	wp.wg.Add(1)
}

// Shutdown shuts down the worker pool.
func (wp *WorkerPool) Shutdown() {
	close(wp.tasks)
	close(wp.shutdown)
	wp.wg.Wait()
}

func main() {
	// Initialize the worker pool with 3 workers
	wp := NewWorkerPool(3)
	wp.Start()

	// Submit sales data tasks
	amazonSales := &SalesData{Marketplace: "Amazon", Revenue: 1000.50}
	ebaySales := &SalesData{Marketplace: "eBay", Revenue: 200.75}
	shopifySales := &SalesData{Marketplace: "Shopify", Revenue: 300.25}

	wp.Submit(amazonSales)
	wp.Submit(ebaySales)
	wp.Submit(shopifySales)

	// Shutdown the worker pool
	wp.Shutdown()

	// Print the total revenue
	fmt.Printf("Total revenue for the day: $%.2f\n", wp.totalRevenue)
}