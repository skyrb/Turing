
package main  
import (  
    "fmt"
    "io/ioutil"
    "os"
    "sync"
    "time"
)

const (
    maxRetries      = 3      // Number of retries for transient errors
    retryDelay      = 100 * time.Millisecond // Delay between retries
    skipUnrecoverable = true // Skip unrecoverable errors
)

func processFile(filePath string, wg *sync.WaitGroup, errors chan error) {
    defer wg.Done()

    // Simulate transient error conditions
    if filePath == "file5.txt" {
        errors <- fmt.Errorf("file is locked for %s", filePath)
        return
    }
    if filePath == "file6.txt" {
        errors <- fmt.Errorf("read timeout for %s", filePath)
        return
    }

    // Simulate unrecoverable error condition
    if filePath == "file7.txt" {
        errors <- fmt.Errorf("file not found: %s", filePath)
        return
    }

    // Process the file with retries
    for i := 0; i <= maxRetries; i++ {
        data, err := ioutil.ReadFile(filePath)
        if err == nil {
            // Process the file content (replace this with your actual processing)
            fmt.Printf("Processed file: %s - Content: %s\n", filePath, string(data))
            return
        }

        if i == maxRetries {
            errors <- err
            return
        }

        fmt.Printf("Transient error occurred for %s, retrying in %v\n", filePath, retryDelay)
        time.Sleep(retryDelay)
    }
}

func main() {
    filePaths := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt", "file7.txt"}
    var wg sync.WaitGroup
    errors := make(chan error, len(filePaths))

    wg.Add(len(filePaths))
    for _, filePath := range filePaths {
        go processFile(filePath, &wg, errors)
    }

    // Wait for all goroutines to finish
    go func() {
        wg.Wait()
        close(errors)
    }()

    // Process errors
    for err := range errors {
        if skipUnrecoverable && os.IsNotExist(err) {
            fmt.Printf("Skipping unrecoverable error: %v\n", err)
            continue
        }
        fmt.Printf("Error: %v\n", err)
    }

    fmt.Println("All files processed.")
} 
  