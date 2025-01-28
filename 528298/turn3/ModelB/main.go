package main  
import (  
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Event represents a civic engagement event.
type Event struct {  
    Name          string `json:"name"`
    ParticipantCount int    `json:"participant_count"`
    SizeCategory  string `json:"-"`
    City          string `json:"city"`
}  

// CityEventStats contains the total participation and size categories for a city.
type CityEventStats struct {  
    TotalParticipation int  
    SmallEvents       int  
    MediumEvents      int  
    LargeEvents       int  
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

// LoadLocations loads location data from a JSON file (assumed to contain City field for each event)
func LoadLocations(filename string) ([]Event, error) {  
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

// CategorizeEvents assigns a size category to each event.  
func CategorizeEvents(events []Event) {
    for i := range events {  
        events[i].SizeCategory = determineSizeCategory(events[i].ParticipantCount)
    }  
}  

// determineSizeCategory returns the category based on the participant count.  
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

// CalculateCityEventStats processes the events to calculate participation and category totals for each city.
func CalculateCityEventStats(events []Event) map[string]CityEventStats {  
    cityStats := make(map[string]CityEventStats)
    for _, event := range events {  
        // Update stats for each city in the event slice
        _, ok := cityStats[event.City]  
        if !ok {
            // Initialize city stats if not already present
            cityStats[event.City] = CityEventStats{}  
        }  
        stats := cityStats[event.City]  
        stats.TotalParticipation += event.ParticipantCount
        switch event.SizeCategory {
        case "Small":
            stats.SmallEvents++  
        case "Medium":
            stats.MediumEvents++  
        case "Large":
            stats.LargeEvents++  
        }  
        cityStats[event.City] = stats  
    }  
    return cityStats  
}  

// FindCityWithHighestParticipation finds the city with the highest total participation.  
func FindCityWithHighestParticipation(cityStats map[string]CityEventStats) (string, int) {  
    highestParticipation := 0
    cityWithHighestParticipation := ""  
    for city, stats := range cityStats {
        if stats.TotalParticipation > highestParticipation {
            highestParticipation = stats.TotalParticipation  
            cityWithHighestParticipation = city
        }  
    }  
    return cityWithHighestParticipation, highestParticipation  
}  

func main() {  
    // Parse command-line arguments  
    eventDataPath := flag.String("events", "events.json", "Path to the events JSON file")