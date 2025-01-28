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

	// Convert to UTC
	utcTime := parsedTime.UTC()
	utcFormatted := formatTime(utcTime, "UTC")

	// Convert to IST
	istTime, err := convertToTimeZone(parsedTime, "Asia/Kolkata")
	if err != nil {
		fmt.Println("Error converting to IST:", err)
		return
	}
	istFormatted := formatTime(istTime, "IST")

	// Print the formatted times
	fmt.Println("UTC:", utcFormatted)
	fmt.Println("IST:", istFormatted)
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