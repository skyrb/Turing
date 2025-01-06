package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// ReadMultipleValues reads data from a file and returns two values: the data and the error.
func ReadMultipleValues(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return data, err
}

func main() {
	filePath := "example.txt" // Replace with your file path

	// Create a channel to receive the result of the read operation
	resultCh := make(chan []byte, 1)
	errCh := make(chan error, 1)

	go func() {
		// Perform the I/O operation in a goroutine
		data, err := ReadMultipleValues(filePath)
		resultCh <- data
		errCh <- err
	}()

	// Wait for the result
	select {
	case data := <-resultCh:
		// Handle data
		fmt.Println("Data received:", string(data))
	case err := <-errCh:
		// Handle error
		log.Fatal("Error reading file:", err)
	}

	// Add a small delay to ensure the program doesn't exit immediately
	time.Sleep(1 * time.Second)
}