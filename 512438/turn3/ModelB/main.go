
package main  
import (  
    "fmt"
    "io/ioutil"
    "os"
    "sync"
    "time"
)
const (  
    retryCount      = 3      // Number of retry attempts for transient errors
    retryInterval   = 1 * time.Second // Interval between retry attempts
    maxErrorCount  = 10     // Maximum number of errors allowed before stopping processing
    errorBufferSize = 1000   // Size of the error channel buffer
)
type fileError struct {
    filePath string
    error    error
}
func processFile(filePath string, wg *sync.WaitGroup, errors chan<- fileError) {
    defer wg.Done()
    var err error
    var data []byte
    // Retry loop for transient errors
    for i := 0; i <= retryCount; i++ {
        data, err = ioutil.ReadFile(filePath)
        if err == nil {
            break
        }
        // Check if the error is transient (file locked)
        if os.IsPermission(err) {
            time.Sleep(retryInterval)
        } else {
            break // Non-transient error, exit the retry loop
        }
    }
    if err != nil {
        // Unrecoverable error, send it to the errors channel
        errors <- fileError{filePath: filePath, error: err}
        return
    }
    // Process the file content (you can replace this with your actual processing)
    fmt.Printf("Processed file: %s - Content: %s\n", filePath, string(data))
}
func logErrors(errors <-chan fileError, logWg *sync.WaitGroup, errorCount *int) {
    defer logWg.Done()
    for errInfo := range errors {
        fmt.Printf("Error processing %s: %v\n", errInfo.filePath, errInfo.error)
        (*errorCount)++
        if *errorCount >= maxErrorCount {
            close(errors) // Close the channel to stop processing if too many errors
        }
    }
}
func main() {  
    filePaths := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
    var wg sync.WaitGroup  
    var logWg sync.WaitGroup  
    errors := make(chan fileError, errorBufferSize)
    errorCount := 0
    logWg.Add(1)
    go logErrors(errors, &logWg, &errorCount)
    wg.Add(len(filePaths))
    for _, filePath := range filePaths {      
        go processFile(filePath, &wg, errors)
    }
    wg.Wait()
    close(errors)
    logWg.Wait()
    if errorCount >= maxErrorCount {
        fmt.Printf("Exiting due to too many errors (%d).\n", errorCount)
    } else {
        fmt.Println("All files processed.")
    }
} 
 