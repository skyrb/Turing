package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// Activity represents a fitness activity
type Activity struct {
	Name        string
	Duration    int
	HeartRate   int
	Time        time.Time
	DayOfWeek   string
}

// AggregatedActivity stores aggregated information for an activity
type AggregatedActivity struct {
	Name          string
	TotalDuration int
	TotalHeartRate int
	Count          int
	AverageHeartRate int
}

// User holds a user's activities
type User struct {
	Name    string
	Activities []Activity
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

// mostCommonTime finds the most common time for high-intensity activities like running
func mostCommonTime(activities []Activity) (time.Time, int) {
	times := make([]time.Time, 0)

	for _, activity := range activities {
		if activity.Name == "Running" && activity.HeartRate > 120 {
			times = append(times, activity.Time)
		}
	}

	if len(times) == 0 {
		return time.Time{}, 0
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Hour() < times[j].Hour()
	})

	maxCount := 1
	currentCount := 1
	mostCommonHour := times[0].Hour()

	for i := 1; i < len(times); i++ {
		if times[i].Hour() == times[i-1].Hour() {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
				mostCommonHour = times[i-1].Hour()
			}
			currentCount = 1
		}
	}

	if currentCount > maxCount {
		mostCommonHour = times[len(times)-1].Hour()
	}

	return time.Time{}.AddDate(0, 0, 0).Add(time.Hour(mostCommonHour)), maxCount
}

func main() {
	// Sample fitness data for multiple users
	users := []User{
		{
			Name: "Alice",
			Activities: []Activity{
				{"Running", 30, 145, time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC), "Sunday"},
				{"Cycling", 45, 130, time.Date(2023, 10, 2, 14, 0, 0, 0, time.UTC), "Monday"},
				{"Running", 25, 140, time.Date(2023, 10, 3, 7, 30, 0, 0, time.UTC), "Tuesday"},
			},
		},
		{
			Name: "Bob",
			Activities: []Activity{
				{"Running", 60, 142, time.Date(2023, 10, 1, 7, 0, 0, 0, time.UTC), "Sunday"},
				{"Swimming", 60, 110, time.Date(2023, 10, 2, 10, 0, 0, 0, time.UTC), "Monday"},
				{"Running", 50, 138, time.Date(2023, 10, 3, 8, 30, 0, 0, time.UTC), "Tuesday"},
			},
		},
	}

	// Flatten the activities for aggregation
	allActivities := []Activity{}
	for _, user := range users {
		allActivities = append(allActivities, user.Activities...)
	}

	// Aggregate activities with heart rate > 120
	aggregatedActivities := aggregateActivities(allActivities, 120)

	fmt.Println("Aggregated activities with heart rate > 120:")
	for _, agg := range aggregatedActivities {
		fmt.Printf("Name: %s, Total Duration: %d minutes, Average Heart Rate: %d bpm\n", agg.Name, agg.TotalDuration, agg.AverageHeartRate)
	}

	// Find the most common time for high-intensity running activities
	mostCommonHour, count := mostCommonTime(allActivities)
	if count > 0 {
		fmt.Printf("Most common time for high-intensity running: %s (occurs %d times)\n", mostCommonHour.Format("15:04"), count)
	} else {
		fmt.Println("No high-intensity running activities found.\n")
	}
}