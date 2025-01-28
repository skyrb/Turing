package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
)

// Event represents a civic engagement event.
type Event struct {
    Name          string `json:"name"`
    ParticipantCount int    `json:"participant_count"`
    SizeCategory  string `json:"-"`
}

// LoadEvents loads event data from a JSON file.
func LoadEvents(filename string) ([]Event, error) {
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var events []Event
    err = json.Unmarshal(file, &events)
    if err != nil {
        return nil, err
    }

    return events, nil
}

// CalculateTotalParticipants sums the participants from all events.
func CalculateTotalParticipants(events []Event) int {
    total := 0
    for _, event := range events {
        total += event.ParticipantCount
    }
    return total
}

// CategorizeEvents assigns a size category to each event.
func CategorizeEvents(events []Event) {
    for i := range events {
        events[i].SizeCategory = determineSizeCategory(events[i].ParticipantCount)
    }
}

// determineSizeCategory returns the category based on participant count.
func determineSizeCategory(participantCount int) string {
    switch {
    case participantCount <= 50:
        return "Small"
    case participantCount <= 200:
        return "Medium"
    default:
        return "Large"
    }
}

func main() {
    // Parse command-line arguments
    datasetPath := flag.String("dataset", "events.json", "Path to the events JSON file")
    flag.Parse()

    // Load events data from the provided JSON file
    events, err := LoadEvents(*datasetPath)
    if err != nil {
        log.Fatalf("Failed to load events data: %v", err)
    }

    // Categorize events into size categories
    CategorizeEvents(events)

    // Calculate the total number of participants
    totalParticipants := CalculateTotalParticipants(events)

    // Output the result
    fmt.Printf("Total participants across all events: %d\n", totalParticipants)
    fmt.Println("Event Categories:")

    for _, event := range events {
        fmt.Printf("Event: %s - Participants: %d - Category: %s\n", event.Name, event.ParticipantCount, event.SizeCategory)
    }
}