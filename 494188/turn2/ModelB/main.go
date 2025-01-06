
package main

import (  
    "fmt"
    "io"
    "log"
    "os"
    "time"
)

// ioResult holds the data and error returned from an I/O operation
type ioResult struct {
    Data []byte
    Err  error
}

// ReadMultipleValues reads data from a file and returns a ioResult.
func ReadMultipleValues(filePath string) ioResult {
    file, err := os.Open(filePath)
    if err != nil {
        return ioResult{Err: err}
    }
    defer file.Close()

    data, err := io.ReadAll(file)    
    return ioResult{Data: data, Err: err}
}

func main() {
    filePaths := []string{"example1.txt", "example2.txt", "example3.txt"} // Replace with your file paths

    resultCh := make(chan ioResult)

    // Spawn a goroutine for each file to perform the I/O operation
    for _, filePath := range filePaths {
        go func(filePath string) {
            result := ReadMultipleValues(filePath)
            resultCh <- result
        }(filePath)
    }

    // Use a loop to process results in the order they are received
    for _, filePath := range filePaths {
        result := <-resultCh
        if result.Err != nil {
            // Priority handling for errors
            log.Printf("Error reading file %q: %v", filePath, result.Err)
            continue
        }

        // Handle data
        fmt.Println("Data received:", string(result.Data))
    }

    // Add a small delay to ensure the program doesn't exit immediately
    time.Sleep(1 * time.Second)
}
 