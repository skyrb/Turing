package main

import (
	"fmt"
	"sync"
)

// SalesData represents sales data for a specific revenue type.
type SalesData struct {
	Marketplace string
	RevenueType string
	Revenue      float64
}

// WorkerPool manages a fixed number of worker goroutines.
type WorkerPool struct {
	tasks     chan *SalesData
	workers   int
	wg        *sync.WaitGroup
	shutdown  chan struct{}
	totalRevenue map[string]map[string]float64
	mu           sync.Mutex
}

// NewWorkerPool creates a new worker pool with the specified number of workers.
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		tasks:     make(chan *SalesData, 100), // Buffered channel
		workers:   workers,
		wg:        &sync.WaitGroup{},
		shutdown:  make(chan struct{}),
		totalRevenue: make(map[string]map[string]float64),
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
	fmt.Printf("Processing %s sales data from %s: $%.2f\n", task.RevenueType, task.Marketplace, task.Revenue)
	wp.mu.Lock()
	if _, ok := wp.totalRevenue[task.Marketplace]; !ok {
		wp.totalRevenue[task.Marketplace] = make(map[string]float64)
	}
	wp.totalRevenue[task.Marketplace][task.RevenueType] += task.Revenue
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

	// Submit sales data tasks for each marketplace and revenue type
	marketplaces := []string{"Amazon", "eBay", "Shopify"}
	revenueTypes := []string{"Product", "Shipping", "Tax"}

	for _, marketplace := range marketplaces {
		for _, revenueType := range revenueTypes {
			revenue := float64(100 + 10*marketplaces.Index(marketplace) + 20*revenueTypes.Index(revenueType))
			task := &SalesData{Marketplace: marketplace, RevenueType: revenueType, Revenue: revenue}
			wp.Submit(task)
		}
	}

	// Shutdown the worker pool
	wp.Shutdown()

	// Print the total revenue for each marketplace and the total revenue across all marketplaces
	fmt.Println("\nTotal Revenue for each Marketplace:")
	for marketplace, revenueData := range wp.totalRevenue {
		fmt.Printf("%s:\n", marketplace)