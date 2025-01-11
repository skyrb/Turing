package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

// processFile reads a given file and simulates errors randomly
func processFile(filePath string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Simulate a random delay
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// Simulate errors with a 50% chance
	if rand.Intn(2) == 0 {
		// Simulate a missing file error
		return os.ErrNotExist
	} else if rand.Intn(2) == 0 {
		// Simulate a permission denial error
		return os.ErrPermission
	}

	// Otherwise, read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("Read file: %s, Content: %s\n", filePath, string(content))
	return nil
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	// Create a channel to hold error messages
	errors := make(chan error, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := processFile(file, &wg)
			if err != nil {
				errors <- err
			}
		}(file)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	close(errors)

	// Collect and print errors
	for err := range errors {
		fmt.Printf("Error processing file: %s\n", err)
	}
}