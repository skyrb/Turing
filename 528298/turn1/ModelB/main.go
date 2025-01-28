
package main  
import (  
    "fmt"
    "encoding/csv"
    "os"
    "strconv"
    "log"
)

type Event struct {
    Name string
    Participants int
}

func main() {
    events := []Event{
        {"Community Cleanup", 20},
        {"Book Club Meeting", 15},
        {"Park Art Exhibition", 12},
        {"Senior Citizen Forum", 25},
    }
    totalParticipants := calculateTotalParticipants(events)
    fmt.Printf("Total participants across all events: %d\n", totalParticipants)
}

func calculateTotalParticipants(events []Event) int {
    total := 0
    for _, event := range events {
        total += event.Participants
    }
    return total
}
