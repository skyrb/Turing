
package main  
import (  
    "fmt"
    "io/ioutil"
    "sync"
    "time"
)

const (
    errorBufferSize = 1000 // Adjust the buffer size as needed
)

func processFile(filePath string, wg *sync.WaitGroup, errors chan<- error) {
    defer wg.Done()

    // Simulate various error conditions
    if filePath == "file2.txt" {
        errors <- fmt.Errorf("permission denied for %s", filePath)
        return
    }
    if filePath == "file4.txt" {
        errors <- fmt.Errorf("file not found: %s", filePath)
        return
    }
    if filePath == "file5.txt" {
        errors <- fmt.Errorf("read timeout for %s", filePath)
        return
    }
    if filePath == "file6.txt" {
        errors <- fmt.Errorf("file is locked for %s", filePath)
        return
    }

    // Simulate file read with a timeout (adjust the duration as needed)
    data, err := readFileWithTimeout(filePath, 2*time.Second)
    if err != nil {
        errors <- err
        return
    }

    // Process the file content (replace this with your actual processing)
    fmt.Printf("Processed file: %s - Content: %s\n", filePath, string(data))
}

func readFileWithTimeout(filePath string, timeout time.Duration) ([]byte, error) {
    c := make(chan []byte, 1)
    go func() {
        data, err := ioutil.ReadFile(filePath)
        c <- data
        if err != nil {
            c <- nil
        }
    }()
    select {
    case data := <-c:
        return data, nil
    case <-time.After(timeout):
        return nil, fmt.Errorf("read timeout for %s", filePath)
    }
}

func logErrors(errors <-chan error, wg *sync.WaitGroup) {
    defer wg.Done()
    for err := range errors {
        fmt.Printf("Error: %v\n", err)
    }
}

func main() {  
    filePaths := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt", "file7.txt", "file8.txt", "file9.txt", "file10.txt"} // Add more file paths as needed
    var wg sync.WaitGroup
    errors := make(chan error, errorBufferSize)

    wg.Add(1)
    go logErrors(errors, &wg)

    wg.Add(len(filePaths))
    for _, filePath := range filePaths {
        go processFile(filePath, &wg, errors)
    }

    // Wait for all goroutines to finish
    wg.Wait()
    close(errors)

    fmt.Println("All files processed.")
} 
