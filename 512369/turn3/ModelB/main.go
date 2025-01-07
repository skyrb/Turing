package main

import (
	"fmt"
	"math"
	"time"
)

// Activity represents a fitness activity
type Activity struct {
	Name        string
	Duration    int
	HeartRate    int
	Time        time.Time
}

// AggregatedActivity stores aggregated information for an activity
type AggregatedActivity struct {
	Name          string
	TotalDuration int
	TotalHeartRate int
	Count          int
	AverageHeartRate int
	MostCommonTime  time.Time // The most common time for high-intensity activities
}

// aggregateActivities aggregates activities by user ID and name, calculating total duration and average heart rate
func aggregateActivities(activities []Activity, threshold int) map[int]map[string]*AggregatedActivity {
	aggregated := make(map[int]map[string]*AggregatedActivity)

	for _, activity := range activities {
		userID := 1 // Replace this with the actual user ID from your data source
		activityName := activity.Name

		if activity.HeartRate > threshold {
			if userAgg, exists := aggregated[userID]; exists {
				if agg, exists := userAgg[activityName]; exists {
					agg.TotalDuration += activity.Duration
					agg.TotalHeartRate += activity.HeartRate
					agg.Count++
					// Update the most common time based on the current activity time
					if agg.Count == 1 {
						agg.MostCommonTime = activity.Time
					} else {
						// Assuming we want the most common time in minutes
						if activity.Time.Minute() == agg.MostCommonTime.Minute() {
							agg.MostCommonTime = activity.Time
						} else {
							agg.MostCommonTime = activity.Time
						}
					}
				} else {
					userAgg[activityName] = &AggregatedActivity{
						Name:          activityName,
						TotalDuration: activity.Duration,
						TotalHeartRate: activity.HeartRate,
						Count:          1,
						MostCommonTime: activity.Time,
					}
				}
			} else {
				aggregated[userID] = make(map[string]*AggregatedActivity)
				aggregated[userID][activityName] = &AggregatedActivity{
					Name:          activityName,
					TotalDuration: activity.Duration,
					TotalHeartRate: activity.HeartRate,
					Count:          1,
					MostCommonTime: activity.Time,
				}
			}
		}
	}

	for _, userAgg := range aggregated {
		for _, agg := range userAgg {
			if agg.Count > 0 {
				agg.AverageHeartRate = int(math.Ceil(float64(agg.TotalHeartRate) / float64(agg.Count)))
			}
		}
	}

	return aggregated
}

func main() {
	// Sample fitness data (Note: Replace this with actual user data)
	activities := []Activity{
		{Name: "Running", Duration: 30, HeartRate: 145, Time: time.Date(2023, 07, 28, 10, 30, 0, 0, time.UTC)},
		{Name: "Cycling", Duration: 45, HeartRate: 130, Time: time.Date(2023, 07, 29, 16, 00, 0, 0, time.UTC)},