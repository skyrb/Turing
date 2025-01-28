package main

import (
    "fmt"
    "time"
)

func main() {
    // Original timestamp string
    timestampStr := "2025-01-25T15:04:05Z"

    // Parse the timestamp string into a time.Time object
    parsedTime, err := time.Parse(time.RFC3339, timestampStr)
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return
    }

    // Format the time into the desired string format
    formattedTime := fmt.Sprintf("%s %d, %d, %02d:%02d %s UTC",
        parsedTime.Month(),
        parsedTime.Day(),
        parsedTime.Year(),
        parsedTime.Hour()%12,
        parsedTime.Minute(),
        formatAMPM(parsedTime.Hour()))

    // Print the formatted time
    fmt.Println(formattedTime)
}

// Helper function to determine AM or PM
func formatAMPM(hour int) string {
    if hour < 12 {
        return "AM"
    }
    return "PM"
}