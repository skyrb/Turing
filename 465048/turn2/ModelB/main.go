package main  
import (  
    "context"
    "fmt"
    "sync"
    "time"
)

func processData(ctx context.Context, data <-chan int, wg *sync.WaitGroup, done chan bool) {  
    defer wg.Done()  

    for {  
        select {  
        case value, ok := <-data:  
            if !ok {  
                return  
            }  
            time.Sleep(time.Duration(value) * time.Millisecond)  
            fmt.Println("Processed data:", value)  
        case <-ctx.Done():  
            fmt.Println("Context canceled, terminating processing.")  
            return  
        }  
    }  
}

func main() {  
    const numGoroutines = 5  
    var wg sync.WaitGroup  
    done := make(chan bool, numGoroutines)

    dataStream1 := make(chan int)
    dataStream2 := make(chan int)
    
    go func() {  
        defer close(dataStream1)  
        for i := 1; i <= 10; i++ {  
            dataStream1 <- i  
        }  
    }()  
    
    go func() {  
        defer close(dataStream2)  
        for i := 11; i <= 20; i++ {  
            dataStream2 <- i  
        }  
    }()  

    // Set a timeout for data processing
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    for i := 0; i < numGoroutines; i++ {  
        wg.Add(1)  
        go processData(ctx, dataStream1, &wg, done)  
        go processData(ctx, dataStream2, &wg, done)  
    }  

    wg.Wait()  
    close(done)  

    if ctx.Err() == context.DeadlineExceeded {  
        fmt.Println("Data processing timeout reached.")  
    }
    fmt.Println("All data streams processed.")  
}