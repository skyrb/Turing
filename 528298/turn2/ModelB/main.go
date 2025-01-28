package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Event represents a civic engagement event with a name and participant count.
type Event struct {
    Name          string `json:"name"`
    ParticipantCount int    `json:"participant_count"`
    Size           string `json:"size"`
}

// Size represents the size of an event
type Size string

const (
    // Event sizes
    Small Size = "small"
    Medium Size = "medium"
    Large  Size = "large"
)

// LoadEvents loads event data from a JSON file.
func LoadEvents(filename string) ([]Event, error) {
    // ... (same code as before)
}

// ProcessEvents applies a callback to each event, updating its size accordingly.
func ProcessEvents(events []Event, categorizer func(Event) Event) []Event {
    for i, event := range events {
        events[i] = categorizer(event)
    }
    return events
}

// CategorizeEventBySize assigns a size to an event based on its participant count.
func CategorizeEventBySize(event Event) Event {
    count := event.ParticipantCount
    switch {
    case count < 100:
        event.Size = string(Small)
    case count < 200:
        event.Size = string(Medium)
    default:
        event.Size = string(Large)
    }
    return event
}

// CalculateTotalParticipants sums up the participants from all events in a given category.
func CalculateTotalParticipantsBySize(events []Event, size Size) int {
    total := 0
    for _, event := range events {
        if event.Size == string(size) {
            total += event.ParticipantCount
        }
    }
    return total
}

func main() {
    datasetPath := flag.String("dataset", "events.json", "Path to the events JSON file")
    flag.Parse()

    // Load events data from the provided JSON file
    events, err := LoadEvents(*datasetPath)
    if err != nil {
        log.Fatalf("Failed to load events data: %v", err)
    }

    // Categorize events based on size
    events = ProcessEvents(events, CategorizeEventBySize)

    // Calculate the total participants for each size
    smallParticipantCount := CalculateTotalParticipantsBySize(events, Small)
    mediumParticipantCount := CalculateTotalParticipantsBySize(events, Medium)
    largeParticipantCount := CalculateTotalParticipantsBySize(events, Large)

    // Output the result
    fmt.Printf("Total participants in small events: %d\n", smallParticipantCount)
    fmt.Printf("Total participants in medium events: %d\n", mediumParticipantCount)
    fmt.Printf("Total participants in large events: %d\n", largeParticipantCount)
}  