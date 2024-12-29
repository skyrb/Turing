package main  
import (  
    "fmt"
    "sync"
)  
func dynamicFormat(names []string) string {  
    var wg sync.WaitGroup  
    maxLength := 0  
    results := make([]string, len(names))  

    for i, name := range names {  
        wg.Add(1)  
        go func(i int, name string) {  
            defer wg.Done()  
            if len(name) > maxLength {  
                maxLength = len(name)  
            }  
            results[i] = fmt.Sprintf(fmt.Sprintf("%%%ds", maxLength), name)  
        }(i, name)  
    }  

    wg.Wait()  
    return fmt.Sprintln(results...)  
}  

func main() {  
    names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry", "Isabella", "Jack"}  
    fmt.Println(dynamicFormat(names))  
} 