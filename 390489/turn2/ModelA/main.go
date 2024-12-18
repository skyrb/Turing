package main

import (
    "fmt"
    "net/http"
    "strings"
    "sync"
    "time"
)

type APICall struct {
    url     string
    timeout time.Duration
    response string
    err     error
}

func makeAPICall(apiCall *APICall) {
    // Make the HTTP GET request
    resp, err := http.Get(apiCall.url)
    if err != nil {
        apiCall.err = err
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        apiCall.err = err
        return
    }
    apiCall.response = string(body)
}

func main() {
    var wg sync.WaitGroup

    // Define multiple API calls with individual timeouts
    calls := []*APICall{
        {
            url:     "https://jsonplaceholder.typicode.com/todos/1",
            timeout: 2 * time.Second,
        },
        {
            url:     "https://jsonplaceholder.typicode.com/todos/2",
            timeout: 1 * time.Second,
        },
        {
            url:     "https://jsonplaceholder.typicode.com/todos/3",
            timeout: 3 * time.Second,
        },
        // Add more calls as needed
    }

    for _, call := range calls {
        // Start each API call in a goroutine
        wg.Add(1)
        go func(apiCall *APICall) {
            defer wg.Done()

            ctx, cancel := context.WithTimeout(context.Background(), apiCall.timeout)
            defer cancel()

            makeAPICall(apiCall)

            select {
            case <-ctx.Done():
                fmt.Println("Context deadline exceeded (timed out) for URL:", apiCall.url)
            default:
                if apiCall.err != nil {
                    fmt.Println("Error for URL:", apiCall.url, ":", apiCall.err)
                } else {
                    fmt.Println("Response for URL:", apiCall.url, ":", apiCall.response)
                }
            }
        }(call)
    }

    // Wait for all goroutines to finish
    wg.Wait()

    fmt.Println("All API calls complete.")
}