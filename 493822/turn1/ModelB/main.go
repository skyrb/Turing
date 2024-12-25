package main  
import (  
    "fmt"
    "time"
)

// Simulate a streaming connection
func simulateStreamingConnection(ch chan<- int) {
    defer close(ch)
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        ch <- i
    }
}

func main() {
    ch := make(chan int)
    go simulateStreamingConnection(ch)

    for data := range ch {
        fmt.Println("Received data:", data)
    }

    // The defer statement will close the channel here, even if there are still
    // pending data in the channel.
    fmt.Println("Streaming connection closed.")
}