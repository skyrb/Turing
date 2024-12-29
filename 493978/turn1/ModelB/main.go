package main  
import (  
    "fmt"
    "net/http"
    "sync"
    "time"
)

var wg sync.WaitGroup

func handleRequests() {  
    http.HandleFunc("/", handler)  
    fmt.Println("Server is listening on port 8080")  
    if err := http.ListenAndServe(":8080", nil); err != nil {  
        panic(err)  
    }  
}  

func handler(w http.ResponseWriter, r *http.Request) {  
    wg.Add(1)  
    defer wg.Done()  
    
    // Simulate some work being done
    time.Sleep(time.Second * 2)
    fmt.Fprintf(w, "Hello World!")  
}  

func main() {  
    go handleRequests()  
    
    // Wait for all goroutines to complete before exiting
    wg.Wait()
}  