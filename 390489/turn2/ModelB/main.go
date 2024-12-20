package main  
import (  
    "fmt"
    "io"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

const (
    apiUrl  = "https://jsonplaceholder.typicode.com/todos/"
    numRequests = 5 // Number of concurrent API calls
)

func simulateApiCall(url string, timeout time.Duration, ch chan<- string) {
    defer close(ch) // Close the channel after the goroutine ends

    // Simulate API call with a random delay
    delay := time.Second + time.Duration(rand.Intn(2))*time.Second
    time.Sleep(delay)

    // Make the HTTP GET request with the specified timeout
    client := &http.Client{Timeout: timeout}
    resp, err := client.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error: %v", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body and send it to the channel
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("Error reading response body: %v", err)
        return
    }
    ch <- string(body)
}

func main() {  
    rand.Seed(time.Now().UnixNano())
    var wg sync.WaitGroup // Create a WaitGroup to manage goroutines
    wg.Add(numRequests) // Increment the WaitGroup counter for each goroutine

    for i := 1; i <= numRequests; i++ {
        url := fmt.Sprintf("%s%d", apiUrl, i)
        timeout := time.Second + time.Duration(rand.Intn(3))*time.Second
        
        go func(url string, timeout time.Duration) {
            defer wg.Done() // Decrement the WaitGroup counter when the goroutine ends
            responseChan := make(chan string)

            // Start the API call simulation concurrently
            go simulateApiCall(url, timeout, responseChan)

            // Use a select statement to handle timeout and response
            select {
            case response := <-responseChan:
                fmt.Printf("Response from %s: %s\n", url, response)
            case <-time.After(timeout):
                fmt.Printf("API call to %s timed out.\n", url)
            }
        }(url, timeout)
    }
    
    wg.Wait() // Wait for all goroutines to finish before exiting
} 