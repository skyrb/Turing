package main
 
import (
	"fmt"
	"sync"
	"time"
	"bytes"
	"strings"
)

const (
	numWorkers = 4 // Adjust this based on your system's CPU core count
)

func formatStringsConcurrently(data []string, wg *sync.WaitGroup, results chan []string) {
	defer wg.Done()
	var buffers []*bytes.Buffer
	for i := 0; i < numWorkers; i++ {
		buffers = append(buffers, new(bytes.Buffer))
	}

	// Partition the data into equal-sized chunks for each worker
	dataSize := len(data)
	chunkSize := dataSize / numWorkers
	startIndex := 0
	for i := 0; i < numWorkers; i++ {
		endIndex := startIndex + chunkSize
		if i == numWorkers-1 {
			endIndex = dataSize
		}
		go formatChunk(data[startIndex:endIndex], buffers[i], results)
		startIndex = endIndex
	}
}

func formatChunk(chunk []string, buffer *bytes.Buffer, results chan []string) {
	// Similar to the previous version, compute the maximum string length in the chunk
	maxLength := 0
	for _, name := range chunk {
		if len(name) > maxLength {
			maxLength = len(name)
		}
	}

	staticPrefix := "Name: "
	staticSuffix := "\n"

	for _, name := range chunk {
		padding := maxLength - len(name)
		paddingStr := strings.Repeat(" ", padding)
		buffer.WriteString(staticPrefix)
		buffer.WriteString(name)
		buffer.WriteString(paddingStr)
		buffer.WriteString(staticSuffix)
	}

	// Send the formatted results of this chunk through the channel
	results <- buffer.String()
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry", "Isabella", "Jack", "Kate", "Liam", "Mia", "Noah", "Olivia", "Parker"}
	// Large dataset example to test concurrency benefits
	for i := 0; i < 1000000; i++ {
		names = append(names, names...)
	}

	startTime := time.Now()
	var wg sync.WaitGroup
	results := make(chan []string, numWorkers)

	wg.Add(1)
	go formatStringsConcurrently(names, &wg, results)
	wg.Wait()
	close(results)

	var finalResult []string
	// Collect the formatted results from all workers
	for result := range results {
		finalResult = append(finalResult, result...)
	}

	fmt.Printf("Time taken: %v\n", time.Since(startTime))
	// Ensure that the order of elements is preserved in the final result
	for _, name := range names {
		fmt.Println(name)
	}
}
