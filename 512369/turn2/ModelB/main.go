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

// ActivitySummary represents a summary of similar activities
type ActivitySummary struct {
	Name           string
	TotalDuration  int
	Count          int
	TotalHeartRate int
}

// filterAndAggregateActivities filters activities based on a heart rate threshold and aggregates similar activities
func filterAndAggregateActivities(activities []Activity, threshold int) []ActivitySummary {
	summaryMap := make(map[string]ActivitySummary)

	for _, activity := range activities {
		if activity.HeartRate > threshold {
			summary, ok := summaryMap[activity.Name]
			if !ok {
				summary = ActivitySummary{
					Name:           activity.Name,
					TotalDuration:  0,
					Count:          0,
					TotalHeartRate: 0,
				}
			}

			summary.TotalDuration += activity.Duration
			summary.TotalHeartRate += activity.HeartRate
			summary.Count++

			summaryMap[activity.Name] = summary
		}
	}

	var summaries []ActivitySummary
	for _, summary := range summaryMap {
		summaries = append(summaries, summary)
	}

	return summaries
}

func main() {
	activities := []Activity{
		{"Running", 30, 145},
		{"Cycling", 45, 130},
		{"Swimming", 60, 110},
		{"Yoga", 45, 85},
		{"Walking", 60, 100},
		{"Hiking", 90, 125},
		{"Running", 45, 150},  // Another running activity to test aggregation
		{"Cycling", 30, 115},  // Another cycling activity to test aggregation
	}

	highHeartRateThreshold := 120
	activitySummaries := filterAndAggregateActivities(activities, highHeartRateThreshold)

	fmt.Println("Aggregated Activities with heart rate > 120:")
	for _, summary := range activitySummaries {
		averageHeartRate := float64(summary.TotalHeartRate) / float64(summary.Count)
		fmt.Printf("Activity: %s, Total Duration: %d minutes, Count: %d, Average Heart Rate: %.2f bpm\n", summary.Name, summary.TotalDuration, summary.Count, averageHeartRate)
	}
}
  