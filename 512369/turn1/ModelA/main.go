package main

import (
	"fmt"
)

// Activity represents a fitness activity
type Activity struct {
	Name    string
	Duration int
	HeartRate int
}

// filterActivities filters activities based on a heart rate threshold
func filterActivities(activities []Activity, threshold int) []Activity {
	var filteredActivities []Activity
	for _, activity := range activities {
		if activity.HeartRate > threshold {
			filteredActivities = append(filteredActivities, activity)
		}
	}
	return filteredActivities
}

func main() {
	// Sample fitness data
	activities := []Activity{
		{"Running", 30, 145},
		{"Cycling", 45, 130},
		{"Swimming", 60, 110},
		{"Yoga", 45, 85},
		{"Walking", 60, 100},
		{"Hiking", 90, 125},
	}

	// Filter activities with heart rate > 120
	highHeartRateActivities := filterActivities(activities, 120)

	fmt.Println("Activities with heart rate > 120:")
	for _, activity := range highHeartRateActivities {
		fmt.Printf("Name: %s, Duration: %d minutes, Heart Rate: %d bpm\n", activity.Name, activity.Duration, activity.HeartRate)
	}
}