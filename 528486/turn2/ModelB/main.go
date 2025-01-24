package main
import (
	"fmt"
	"time"
)

// Function to count the votes and identify peak engagement times for each option
func analyzeActivityTrends(surveyData map[string]map[string]string) (map[string]int, map[string]time.Time) {
	optionCounts := make(map[string]int)
	peakEngagementTimes := make(map[string]time.Time)

	for _, responses := range surveyData {
		for option, timestamp := range responses {
			optionCounts[option]++

			t, err := time.Parse("2006-01-02 15:04", timestamp)
			if err != nil {
				panic(err)
			}

			currentPeak, exists := peakEngagementTimes[option]
			if !exists || t.After(currentPeak) {
				peakEngagementTimes[option] = t
			}
		}
	}
	return optionCounts, peakEngagementTimes
}

func main() {
	surveyData := map[string]map[string]string{
		"User1": {
			"Option A": "2023-07-31 10:00",
			"Option B": "2023-07-31 10:30",
		},
		"User2": {
			"Option B": "2023-07-31 11:00",
			"Option C": "2023-07-31 11:15",
		},
		"User3": {
			"Option A": "2023-07-31 14:00",
			"Option C": "2023-07-31 14:15",
			"Option B": "2023-07-31 14:30",
		},
		"User4": {
			"Option A": "2023-07-31 17:00",
		},
	}

	voteCounts, peakTimes := analyzeActivityTrends(surveyData)

	fmt.Println("Vote counts for each option:")
	for option, count := range voteCounts {
		fmt.Printf("%s: %d\n", option, count)
	}

	fmt.Println("\nPeak engagement times for each option:")
	for option, peakTime := range peakTimes {
		fmt.Printf("%s: %s\n", option, peakTime.Format("3:04 PM"))
	} 
}  