package main  
import (  
    "fmt"
)  
func main() {  
    names := []string{"Alice", "Bob", "Charlie", "David"}  
    maxLength := 0  
    for _, name := range names {  
        if len(name) > maxLength {  
            maxLength = len(name)  
        }  
    }  
    format := fmt.Sprintf("%%%ds", maxLength)  
    for _, name := range names {  
        fmt.Println(fmt.Sprintf(format, name))  
    }  
}