package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxWorkers = 5
	bufferSize = 100
	dataRate   = time.Millisecond * 50
)

var (
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
	workChan    = make(chan int, bufferSize)
	done        = make(chan struct{})
)

// Produce data until told to stop and wait for all producers to complete
func produceData() {
	defer wgProducers.Done()
	for i := 0; ; i++ {
		select {
		case workChan <- i:
			time.Sleep(dataRate)
		default:
			// Channel is full, backpressure is applied
			fmt.Println("Producer backing off due to full channel.")
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// Consume data from the channel
func consumeData() {
	defer wgConsumers.Done()
	for data := range workChan {
		fmt.Printf("Consuming data: %d\n", data)
		time.Sleep(time.Millisecond * 20) // Simulate processing time
	}
}

func main() {
	// Add producers to producer wait group
	wgProducers.Add(maxWorkers)

	// Start producers
	for i := 0; i < maxWorkers; i++ {
		go produceData()
	}

	// Add consumers to consumer wait group
	wgConsumers.Add(maxWorkers)

	// Start consumers
	for i := 0; i < maxWorkers; i++ {
		go consumeData()
	}

	// Simulate running the system for a while
	time.Sleep(time.Second * 5)

	// Signal producers to stop by waiting for them to finish
	fmt.Println("Stopping producers.")
	wgProducers.Wait()

	// Now that all producers are done, close the work channel
	close(workChan)

	// Wait for all consumers to complete processing
	wgConsumers.Wait()

	fmt.Println("All producers and consumers have stopped. Program exiting gracefully.")
}