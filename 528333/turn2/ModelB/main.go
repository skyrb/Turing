package main  
import (  
    "fmt"
    "time"
)

func main() {  
    timestampStr := "2025-01-25T15:04:05Z"
    parsedTime, err := time.Parse(time.RFC3339, timestampStr)  
    if err != nil {    
        fmt.Println("Error parsing time:", err)    
        return  
    }  
    formattedTimeUTC := formatTime(parsedTime, "UTC")  
    formattedTimeIST := formatTime(parsedTime, "Asia/Kolkata")  
    fmt.Println("UTC:", formattedTimeUTC)  
    fmt.Println("IST:", formattedTimeIST)
}  
func formatTime(t time.Time, location string) string {  
    // Load the time zone location from the IANA database
    loc, err := time.LoadLocation(location)  
    if err != nil {    
        // Handle error if the time zone data is missing or inconsistent
        return fmt.Sprintf("Error loading time zone: %v", err)  
    }  
    // Convert the time to the specified location
    t = t.In(loc)  
    return fmt.Sprintf("%s %d, %d, %02d:%02d %s %s",    
        t.Month(),    
        t.Day(),    
        t.Year(),    
        t.Hour()%12,    
        t.Minute(),    
        formatAMPM(t.Hour()),    
        location,  
    )
}  
func formatAMPM(hour int) string {  
    if hour < 12 {    
        return "AM"  
    }  
    return "PM"
}