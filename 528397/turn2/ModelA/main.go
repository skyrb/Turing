package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a content distribution task with a priority level
type Task struct {
	region   string
	priority int // Lower number indicates higher priority
}

// PriorityQueue is a simple implementation of a priority queue using channels
type PriorityQueue struct {
	tasks chan Task
}

func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{
		tasks: make(chan Task, size),
	}
}

func (pq *PriorityQueue) AddTask(task Task) {
	pq.tasks <- task
}

func (pq *PriorityQueue) GetTask() Task {
	return <-pq.tasks
}

func (pq *PriorityQueue) Close() {
	close(pq.tasks)
}

func distributeContent(task Task, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting content distribution to %s region with priority %d\n", task.region, task.priority)

	// Simulating product image distribution
	distributeProductImages(task.region)

	// Simulating promotional banners distribution
	distributePromotionalBanners(task.region)

	fmt.Printf("Completed content distribution to %s region\n", task.region)
}

func distributeProductImages(region string) {
	fmt.Printf("Distributing product images to %s region...\n", region)
	time.Sleep(2 * time.Second)
	fmt.Printf("Product images distributed to %s region\n", region)
}

func distributePromotionalBanners(region string) {
	fmt.Printf("Distributing promotional banners to %s region...\n", region)
	time.Sleep(2 * time.Second)
	fmt.Printf("Promotional banners distributed to %s region\n", region)
}

func worker(id int, pq *PriorityQueue, wg *sync.WaitGroup) {
	for task := range pq.tasks {
		distributeContent(task, wg)
	}
	fmt.Printf("Worker %d exited\n", id)
}

func main() {
	var wg sync.WaitGroup
	pq := NewPriorityQueue(10)

	// Define the regions and their priorities
	tasks := []Task{
		{"North America", 2},
		{"Europe", 3},
		{"Asia", 1}, // Prioritize Asia due to Singles' Day promotions
	}

	// Add tasks to the priority queue
	for _, task := range tasks {
		wg.Add(1)
		go pq.AddTask(task)
	}

	// Start a pool of worker goroutines
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		go worker(i, pq, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()
	pq.Close()

	fmt.Println("All content distribution tasks completed.")
}