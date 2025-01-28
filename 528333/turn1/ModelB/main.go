package main  
import (  
    "fmt"
    "time"
)

func main() {  
    timestamp := "2025-01-25T15:04:05Z"
    t, err := time.Parse(time.RFC3339, timestamp)  
    if err != nil {
        fmt.Println("Error parsing timestamp:", err)
        return
    }
    formattedTime := fmt.Sprintf("%s %d, %d, %02d:%02d %s UTC", 
        t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute(), 
        t.Format("PM"))
    fmt.Println("Formatted Timestamp:", formattedTime)  
}