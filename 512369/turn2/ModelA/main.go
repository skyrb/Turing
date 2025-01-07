package main

import (
	"fmt"
	"math"
)

// Activity represents a fitness activity
type Activity struct {
	Name    string
	Duration int
	HeartRate int
}

// aggregateActivities aggregates activities by name, calculating total duration and average heart rate
func aggregateActivities(activities []Activity, threshold int) map[string]*AggregatedActivity {
	aggregated := make(map[string]*AggregatedActivity)

	for _, activity := range activities {
		if activity.HeartRate > threshold {
			if agg, exists := aggregated[activity.Name]; exists {
				agg.TotalDuration += activity.Duration
				agg.TotalHeartRate += activity.HeartRate
				agg.Count++
			} else {
				aggregated[activity.Name] = &AggregatedActivity{
					Name:          activity.Name,
					TotalDuration: activity.Duration,
					TotalHeartRate: activity.HeartRate,
					Count:          1,
				}
			}
		}
	}

	for _, agg := range aggregated {
		if agg.Count > 0 {
			agg.AverageHeartRate = int(math.Ceil(float64(agg.TotalHeartRate) / float64(agg.Count)))
		}
	}

	return aggregated
}

// AggregatedActivity stores aggregated information for an activity
type AggregatedActivity struct {
	Name          string
	TotalDuration int
	TotalHeartRate int
	Count          int
	AverageHeartRate int
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
		{"Running", 25, 140},
		{"Cycling", 50, 127},
		{"Hiking", 85, 122},
	}

	// Aggregate activities with heart rate > 120
	aggregatedActivities := aggregateActivities(activities, 120)

	fmt.Println("Aggregated activities with heart rate > 120:")
	for _, agg := range aggregatedActivities {
		fmt.Printf("Name: %s, Total Duration: %d minutes, Average Heart Rate: %d bpm\n", agg.Name, agg.TotalDuration, agg.AverageHeartRate)
	}
}