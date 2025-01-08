package main

import (
	"fmt"
	"sync"
)

// SalesData represents sales data from a single marketplace.
type SalesData struct {
	Marketplace string
	ProductRevenue float64
	ShippingRevenue float64
	TaxRevenue     float64
}

// WorkerPool manages a fixed number of worker goroutines for processing a specific type of revenue.
type WorkerPool struct {
	tasks     chan *SalesData
	workers   int
	wg        *sync.WaitGroup
	shutdown  chan struct{}
	totalRevenue float64
	mu           sync.Mutex
}

// NewWorkerPool creates a new worker pool with the specified number of workers for processing a specific type of revenue.
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

// processTask processes a sales data task for a specific type of revenue and accumulates the total revenue.
func (wp *WorkerPool) processTask(task *SalesData) {
	fmt.Printf("Processing %s revenue from %s: $%.2f\n", task.Marketplace, wp.wg.Name, task.ProductRevenue+task.ShippingRevenue+task.TaxRevenue)
	wp.mu.Lock()
	wp.totalRevenue += task.ProductRevenue + task.ShippingRevenue + task.TaxRevenue
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
	// Initialize worker pools for each type of revenue
	productWP := NewWorkerPool(3)
	shippingWP := NewWorkerPool(3)
	taxWP := NewWorkerPool(3)

	// Start all worker pools
	productWP.Start()
	shippingWP.Start()
	taxWP.Start()

	// Submit sales data tasks
	amazonSales := &SalesData{Marketplace: "Amazon", ProductRevenue: 1000.50, ShippingRevenue: 50.25, TaxRevenue: 100.50}
	ebaySales := &SalesData{Marketplace: "eBay", ProductRevenue: 200.75, ShippingRevenue: 10.50, TaxRevenue: 25.25}
	shopifySales := &SalesData{Marketplace: "Shopify", ProductRevenue: 300.25, ShippingRevenue: 20.75, TaxRevenue: 30.00}

	productWP.Submit(amazonSales)
	productWP.Submit(ebaySales)
	productWP.Submit(shopifySales)

	shippingWP.Submit(amazonSales)
	shippingWP.Submit(ebaySales)
	shippingWP.Submit(shopifySales)

	taxWP.Submit(amazonSales)
	taxWP.Submit(ebaySales)
	taxWP.Submit(shopifySales)

	// Shutdown all worker pools
	productWP.Shutdown()
	shippingWP.Shutdown()
	taxWP.Shutdown()

	// Print total revenue for each type of revenue and for each marketplace
	var totalProductRevenue, totalShippingRevenue, totalTaxRevenue float64
	var amazonTotal, ebayTotal, shopifyTotal float64

	fmt.Println("\nTotal Revenue for Each Marketplace:")
	for marketplace := range map[string]bool{"Amazon": true, "eBay": true, "Shopify": true} {
		amazonTotal += amazonSales.ProductRevenue + amazonSales.ShippingRevenue + amazonSales.TaxRevenue
		ebayTotal += ebaySales.ProductRevenue + ebaySales.ShippingRevenue + ebaySales.TaxRevenue
		shopifyTotal += shopifySales.ProductRevenue + shopifySales.ShippingRevenue + shopifySales.TaxRevenue
	}
	fmt.Printf("Amazon: $%.2f\n", amazonTotal)
	fmt.Printf("eBay: $%.2f\n", ebayTotal)
	fmt.Printf("Shopify: $%.2f\n", shopifyTotal)

	fmt.Println("\nTotal Revenue for Each Type of Revenue:")
	fmt.Printf("Product: $%.2f\n", productWP.totalRevenue)
	fmt.Printf("Shipping: $%.2f\n", shippingWP.totalRevenue)
	fmt.Printf("Tax: $%.2f\n", taxWP.totalRevenue)

	fmt.Println("\nTotal Revenue Across All Marketplaces:")
	fmt.Printf("Total Revenue: $%.2f\n", productWP.totalRevenue+shippingWP.totalRevenue+taxWP.totalRevenue)
}