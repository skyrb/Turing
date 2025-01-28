
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Original timestamp strings in a batch
	timestampStrings := []string{
		"2025-01-25T15:04:05Z",
		"2025-01-26T08:37:12Z",
		"2025-01-27T17:18:56Z",
		// Add more timestamps as needed
	}

	// Batch size for concurrent processing
	batchSize := 100

	// Results will be stored in this slice
	formattedTimes := make([]string, len(timestampStrings))

	// WaitGroup to synchronize concurrent execution
	var wg sync.WaitGroup

	// Process timestamps in batches
	for i := 0; i < len(timestampStrings); i += batchSize {
		end := min(i+batchSize, len(timestampStrings))
		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				formattedTimes[j], _ = processTimestamp(timestampStrings[j])
			}
		}(i, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the formatted times
	for _, formattedTime := range formattedTimes {
		fmt.Println(formattedTime)
	}
}

func processTimestamp(timestampStr string) (string, error) {
	// Parse the timestamp string
	parsedTime, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return "", err
	}

	// Convert to UTC
	utcTime := parsedTime.UTC()
	utcFormatted := formatTime(utcTime, "UTC")

	// Convert to IST
	istTime, err := convertToTimeZone(parsedTime, "Asia/Kolkata")
	if err != nil {
		return "", err
	}
	istFormatted := formatTime(istTime, "IST")

	return fmt.Sprintf("%s | %s", utcFormatted, istFormatted), nil
}

func convertToTimeZone(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} 
  