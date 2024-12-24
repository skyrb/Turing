package main

import (
	"fmt"
	"sync"
	"time"
)

func processData(data <-chan int, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	for value := range data {
		time.Sleep(time.Duration(value) * time.Millisecond)
		fmt.Println("Processed data:", value)
	}

	done <- true
}

func main() {
	const numGoroutines = 5
	var wg sync.WaitGroup
	done := make(chan bool, numGoroutines)

	dataStream1 := make(chan int)
	dataStream2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			dataStream1 <- i
		}
		close(dataStream1)
	}()

	go func() {
		for i := 11; i <= 20; i++ {
			dataStream2 <- i
		}
		close(dataStream2)
	}()

	// Start Goroutines to process the data
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processData(dataStream1, &wg, done)
		go processData(dataStream2, &wg, done)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	close(done)

	fmt.Println("All data streams processed.")
}