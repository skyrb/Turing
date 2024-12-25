package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxGoroutines = 10
	minGoroutines = 2
)

var wg sync.WaitGroup

func processData(ctx context.Context, data <-chan int, done chan bool) {
	defer wg.Done()

	for {
		select {
		case value, ok := <-data:
			if !ok {
				return
			}
			time.Sleep(time.Duration(value) * time.Millisecond)
			fmt.Println("Processed data:", value)
		case <-ctx.Done():
			fmt.Println("Context canceled, terminating processing.")
			return
		}
	}
}

func goroutineManager(ctx context.Context, data <-chan int, load chan int) {
	currentGoroutines := minGoroutines

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine manager terminated.")
			return
		case load := <-load:
			// Adjust the number of Goroutines based on the load
			newGoroutines := adjustGoroutines(currentGoroutines, load)

			if newGoroutines > currentGoroutines {
				// Start new Goroutines
				for i := currentGoroutines; i < newGoroutines; i++ {
					wg.Add(1)
					go processData(ctx, data, nil)
				}
			} else if newGoroutines < currentGoroutines {
				// Cancel existing Goroutines
				for i := currentGoroutines - 1; i >= newGoroutines; i-- {
					go func() {
						time.Sleep(time.Millisecond * 100) // Allow some time for existing work to finish
						done <- true
					}()
				}
			}

			currentGoroutines = newGoroutines
			fmt.Printf("Current Goroutines: %d\n", currentGoroutines)
		}
	}
}

func adjustGoroutines(currentGoroutines int, load int) int {
	// Simple load-based adjustment logic
	if load < 30 {
		return minGoroutines
	} else if load > 70 {
		return maxGoroutines
	}
	return currentGoroutines
}

func main() {
	dataStream := make(chan int)
	loadChan := make(chan int)
	done := make(chan bool, maxGoroutines)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go goroutineManager(ctx, dataStream, loadChan)

	// Simulate data processing and load changes
	go func() {
		for {
			load := rand.Intn(100)
			loadChan <- load
			fmt.Printf("Load: %d\n", load)

			for i := 1; i <= 10; i++ {
				dataStream <- i
			}

			time.Sleep(time.Second * 2)
		}
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
	close(done)

	fmt.Println("All data streams processed.")
}