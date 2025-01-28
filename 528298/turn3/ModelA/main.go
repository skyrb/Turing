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
	Name           string `json:"name"`
	ParticipantCount int    `json:"participant_count"`
	LocationID     string `json:"location_id"`
}

// Location represents the city and other details of an event location.
type Location struct {
	ID   string `json:"id"`
	City string `json:"city"`
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

// LoadLocations loads location data from a JSON file.
func LoadLocations(filename string) (map[string]Location, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var locations []Location
	err = json.Unmarshal(file, &locations)
	if err != nil {
		return nil, err
	}

	locationMap := make(map[string]Location)
	for _, location := range locations {
		locationMap[location.ID] = location
	}

	return locationMap, nil
}

// CalculateParticipantsByCity accumulates participant counts by city.
func CalculateParticipantsByCity(events []Event, locations map[string]Location) map[string]int {
	cityParticipants := make(map[string]int)

	for _, event := range events {
		if location, ok := locations[event.LocationID]; ok {
			cityParticipants[location.City] += event.ParticipantCount
		}
	}

	return cityParticipants
}

func main() {
	// Parse command-line arguments
	eventsPath := flag.String("events", "events.json", "Path to the events JSON file")
	locationsPath := flag.String("locations", "locations.json", "Path to the locations JSON file")
	flag.Parse()

	// Load events data
	events, err := LoadEvents(*eventsPath)
	if err != nil {
		log.Fatalf("Failed to load events data: %v", err)
	}

	// Load locations data
	locations, err := LoadLocations(*locationsPath)
	if err != nil {
		log.Fatalf("Failed to load locations data: %v", err)
	}

	// Calculate participants by city
	cityParticipants := CalculateParticipantsByCity(events, locations)

	// Find the city with the highest cumulative participation
	var maxCity string
	var maxParticipants int
	for city, participants := range cityParticipants {
		if participants > maxParticipants {
			maxCity = city
			maxParticipants = participants
		}
	}

	// Output the result
	fmt.Printf("City with the highest cumulative event participation: %s (%d participants)\n", maxCity, maxParticipants)
}