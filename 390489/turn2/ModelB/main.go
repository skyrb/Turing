package main  
import (  
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

const (
    apiUrl1 = "https://jsonplaceholder.typicode.com/todos/1" 
    apiUrl2 = "https://jsonplaceholder.typicode.com/todos/2" 
    // Add more API URLs here
)

// API Request struct with its own timeout
type apiRequest struct {
    url     string
    timeout time.Duration
}

func simulateApiCall(req apiRequest, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    // Simulate API call with a random delay within the request timeout
    delay := time.Duration(time.Second) + time.Duration(rand.Intn(int(req.timeout)))*time.Second
    time.Sleep(delay)
    
    resp, err := http.Get(req.url)
    if err != nil {
        ch <- fmt.Sprintf("Error for %s: %v", req.url, err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("Error reading response body for %s: %v", req.url, err)
        return
    }
    ch <- fmt.Sprintf("Response for %s: %s", req.url, string(body))
}

func main() {
    // Create a channel to receive responses
    responseChan := make(chan string)

    // Set up waitgroup to track the number of running goroutines
    var wg sync.WaitGroup
    
    // Define the API requests with individual timeouts
    requests := []apiRequest{
        {apiUrl1, 3 * time.Second}, 
        {apiUrl2, 2 * time.Second}, 
        // Add more requests with timeouts here
    }
   
    for _, req := range requests {
        // Add 1 to the waitgroup for each goroutine starting
        wg.Add(1)
        go simulateApiCall(req, responseChan, &wg)
    }

    // Start receiving responses from the channel until all goroutines are complete
    go func() {
        for response := range responseChan {
            fmt.Println(response)
        }
    }()
    
    // Wait for all goroutines to finish their tasks
    wg.Wait()
    
    close(responseChan) // Close the channel to signal completion to the receiver
    fmt.Println("All API calls finished.")
} 