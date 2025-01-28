package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
)

// Event represents a civic engagement event with a name and participant count.
type Event struct {
    Name          string `json:"name"`
    ParticipantCount int    `json:"participant_count"`
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

// CalculateTotalParticipants sums up the participants from all events.
func CalculateTotalParticipants(events []Event) int {
    total := 0
    for _, event := range events {
        total += event.ParticipantCount
    }
    return total
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

    // Calculate the total number of participants
    totalParticipants := CalculateTotalParticipants(events)

    // Output the result
    fmt.Printf("Total participants across all events: %d\n", totalParticipants)
}