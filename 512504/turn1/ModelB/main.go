package main
import (
    "fmt"
    "sync"
    "time"
)
type userProfile struct {
    Name     string
    Age      int
    Location string
}
type userData struct {
    Profile userProfile
    History map[string][]change
}
type change struct {
    Value    interface{}
    Timestamp time.Time
}
func newUserData(profile userProfile) *userData {
    return &userData{
        Profile: profile,
        History: map[string][]change{
            "Name":     {change{profile.Name, time.Now()}},
            "Age":      {change{profile.Age, time.Now()}},
            "Location": {change{profile.Location, time.Now()}},
        },
    }
}
func (ud *userData) updateField(field string, newValue interface{}) {
    ud.mu.Lock()
    defer ud.mu.Unlock()
    switch field {
    case "Name":
        ud.Profile.Name = newValue.(string)
    case "Age":
        ud.Profile.Age = newValue.(int)
    case "Location":
        ud.Profile.Location = newValue.(string)
    default:
        return
    }
    ud.History[field] = append(ud.History[field], change{newValue, time.Now()})
}
func (ud *userData) getFieldHistory(field string) []change {
    ud.mu.RLock()
    defer ud.mu.RUnlock()
    return ud.History[field]
}
var (
    users sync.Map
)
func main() {
    // Initialize user data
    users.Store(1, newUserData(userProfile{"Alice", 25, "New York"}))
    users.Store(2, newUserData(userProfile{"Bob", 30, "San Francisco"}))
    // Concurrent updates
    go func() {
        users.LoadAndUpdate(1, func(value interface{}) interface{} {
            ud := value.(*userData)
            ud.updateField("Age", 26)
            ud.updateField("Location", "Los Angeles")
            return ud
        })
    }()
    go func() {
        users.LoadAndUpdate(2, func(value interface{}) interface{} {
            ud := value.(*userData)
            ud.updateField("Name", "Robert")
            return ud
        })
    }()
    // Wait for updates to complete
    time.Sleep(time.Second)
    // Print user profiles and histories
    users.Range(func(key, value interface{}) bool {
        ud := value.(*userData)
        fmt.Printf("User %d: %+v\n", key, ud.Profile)
        fmt.Println("History:")
        for field, history := range ud.History {
            fmt.Printf("  %s: %+v\n", field, history)
        }
        return true
    })
}  