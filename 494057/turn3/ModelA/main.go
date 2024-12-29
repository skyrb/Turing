package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const workerCount = 4

func formatStringChunk(dataChunk []string, maxLength int, results chan<- string, wg *sync.WaitGroup) {
	var buffer bytes.Buffer
	staticPrefix := "Name: "
	staticSuffix := "\n"

	defer wg.Done()

	for _, name := range dataChunk {
		// Calculate padding required for the string
		padding := maxLength - len(name)
		paddingStr := strings.Repeat(" ", padding)

		// Write formatted string to the buffer
		buffer.WriteString(staticPrefix)
		buffer.WriteString(name)
		buffer.WriteString(paddingStr)
		buffer.WriteString(staticSuffix)
	}

	results <- buffer.String()
}

func formatStringsConcurrently(data []string) string {
	var results []string
	var maxLength int
	var wg sync.WaitGroup
	resultsCh := make(chan string, workerCount)

	// Determine the maximum string length
	for _, name := range data {
		if len(name) > maxLength {
			maxLength = len(name)
		}
	}

	// Partition data into chunks
	chunkSize := len(data) / workerCount
	if len(data)%workerCount > 0 {
		chunkSize++
	}

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go formatStringChunk(data[i:end], maxLength, resultsCh, &wg)
	}

	// Collect results
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// Gather formatted strings from the channel
	for result := range resultsCh {
		results = append(results, result)
	}

	// Combine the results into one final string
	return strings.Join(results, "")
}

func main() {
	names := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		names[i] = fmt.Sprintf("Name %06d", i)
	}

	start := time.Now()
	formatted := formatStringsConcurrently(names)
	end := time.Now()

	fmt.Printf("Formatted output:\n%s\n", formatted)
	fmt.Printf("Formatted %d names in %s\n", len(names), end.Sub(start))
}