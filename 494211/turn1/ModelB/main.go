package main  
import (  
    "fmt"
    "strings"
    "sync"
)

//FuncName - my concurrent function name 
func FuncName(s string, wg *sync.WaitGroup, ch chan []string) {  
    defer wg.Done() // Without this we'll get deadlock  

    s1 := strings.ToUpper(s)  
    s2 := strings.ToLower(s)  
    s3 := strings.Split(s, " ")  
    res := []string{s1, s2, strings.Join(s3, "-")}
    ch <- res  
}  

func main() {  
    strSlice := []string{"Go", "Concurrent", "Programming"}  
    ch := make(chan []string)  
    var wg sync.WaitGroup  

    // Use range to iterate over the strSlice
    for _, s := range strSlice {  
        wg.Add(1)  
        go FuncName(s, &wg, ch)  
    }  
    wg.Wait()  

    // Results collection
    results := make([]string, 0)

    // Close the channel
    close(ch)

    // Use for loop to receive all the results from the channel.
    for res := range ch {
        results = append(results, res...)
    }
    
    fmt.Println(results)  
}