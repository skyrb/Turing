package main  
import (  
    "fmt"
    "io/ioutil"
    "sync"
)

func processFile(filePath string, wg *sync.WaitGroup, errors chan error) {
    defer wg.Done()

    // Simulate error conditions
    if filePath == "file2.txt" {
        errors <- fmt.Errorf("permission denied for %s", filePath)
        return
    }
    if filePath == "file4.txt" {
        errors <- fmt.Errorf("file not found: %s", filePath)
        return
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

func main() {
    filePaths := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}
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
        fmt.Printf("Error: %v\n", err)
    }

    fmt.Println("All files processed.")
} 