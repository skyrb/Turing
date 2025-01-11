package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
	"math/rand"
)

func processFile(filePath string, ctx context.Context, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()

	// Simulate potential errors
	select {
	case <-ctx.Done():
		errors <- fmt.Errorf("context canceled for %s", filePath)
	case <-time.After(time.Duration(rand.Intn(3)) * time.Second):
		if filePath == "file2.txt" {
			errors <- fmt.Errorf("permission denied for %s", filePath)
			return
		}
		if filePath == "file4.txt" {
			errors <- fmt.Errorf("file not found: %s", filePath)
			return
		}
		if rand.Intn(10) == 0 {
			errors <- fmt.Errorf("read timeout for %s", filePath)
			return
		}
	}

	// Read the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		errors <- err
		return
	}

	// Process the file content (you can replace this with your actual processing)
	fmt.Printf("Processed file: %s - Content: %s\n", filePath, string(data))
}

func collectErrors(errors <-chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for err := range errors {
		log.Printf("Error: %v\n", err)
	}
}

func main() {
	filePaths := make([]string, 0, 100)
	for i := 1; i <= 100; i++ {
		filePaths = append(filePaths, fmt.Sprintf("file%d.txt", i))
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(filePaths))

	// Start a goroutine to collect errors
	wg.Add(1)
	go collectErrors(errors, &wg)

	// Limit the number of concurrent goroutines
	const maxConcurrency = 10
	var semaphore = make(chan struct{}, maxConcurrency)

	wg.Add(len(filePaths))
	for _, filePath := range filePaths {
		go func(file string) {
			defer func() {
				semaphore <- struct{}{}
			}()
			<-semaphore // Request a slot from the semaphore

			// Use context to handle cancellations
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			processFile(file, ctx, &wg, errors)
		}(filePath)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errors)

	fmt.Println("All files processed.")
}