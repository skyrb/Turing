package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    apiUrl    = "https://jsonplaceholder.typicode.com/todos/1"
    timeout   = 2 * time.Second // Timeout duration for API call
)

func simulateApiCall(url string, ch chan<- string) {
    // Simulate API call with a random delay
    delay := time.Duration(time.Second + time.Duration(rand.Intn(2))*time.Second)
    time.Sleep(delay)

    // Make the HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error: %v", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body and send it to the channel
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("Error reading response body: %v", err)
        return
    }
    ch <- string(body)
}

func main() {
    // Create a channel to receive responses
    responseChan := make(chan string)

    // Start the API call simulation concurrently
    go simulateApiCall(apiUrl, responseChan)

    // Use a select statement to handle timeout and response
    select {
    case response := <-responseChan:
        fmt.Println("Response:", response)
    case <-time.After(timeout):
        fmt.Println("API call timed out.")
    }
}