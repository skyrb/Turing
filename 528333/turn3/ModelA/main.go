package main

import (
	"fmt"
	"sync"
	"time"
)

// Timestamp struct represents a log entry with a timestamp
type Timestamp struct {
	ID        int
	Timestamp string
}

func main() {
	// Example list of log entries with timestamps
	timestamps := []Timestamp{
		{ID: 1, Timestamp: "2025-01-25T15:04:05Z"},
		{ID: 2, Timestamp: "2023-03-12T12:34:56Z"},
		{ID: 3, Timestamp: "2024-10-30T08:00:00Z"},
	}

	// Channel to collect formatted results
	resultsChan := make(chan string, len(timestamps))

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Process each timestamp concurrently
	for _, entry := range timestamps {
		wg.Add(1)
		go func(entry Timestamp) {
			defer wg.Done()

			// Parse the input timestamp
			parsedTime, err := time.Parse(time.RFC3339, entry.Timestamp)
			if err != nil {
				resultsChan <- fmt.Sprintf("Error parsing time for ID %d: %v", entry.ID, err)
				return
			}

			// Format the time into both UTC and IST
			utcFormatted := formatTime(parsedTime.UTC(), "UTC")
			istTime, err := convertToTimeZone(parsedTime, "Asia/Kolkata")
			if err != nil {
				resultsChan <- fmt.Sprintf("Error converting to IST for ID %d: %v", entry.ID, err)
				return
			}
			istFormatted := formatTime(istTime, "IST")

			// Send the formatted string to the results channel
			resultsChan <- fmt.Sprintf("ID %d: UTC: %s, IST: %s", entry.ID, utcFormatted, istFormatted)
		}(entry)
	}

	// Close the results channel once all processing is done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect and print the results
	for result := range resultsChan {
		fmt.Println(result)
	}
}

// Helper function to convert time to a specified time zone
func convertToTimeZone(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

// Helper function to format time in a specific format with time zone abbreviation
func formatTime(t time.Time, tzAbbr string) string {
	return fmt.Sprintf("%s %d, %d, %02d:%02d %s %s",
		t.Month(),
		t.Day(),
		t.Year(),
		hourTo12Hour(t.Hour()),
		t.Minute(),
		formatAMPM(t.Hour()),
		tzAbbr)
}

// Convert 24-hour format to 12-hour format
func hourTo12Hour(hour int) int {
	h := hour % 12
	if h == 0 {
		h = 12
	}
	return h
}

// Determine AM or PM
func formatAMPM(hour int) string {
	if hour < 12 {
		return "AM"
	}
	return "PM"
}