package main  
import (  
    "fmt"
    "time"
)

// UserHealthData represents health data associated with a user.
type UserHealthData struct {  
    HeartRate    int
    ActivityLevel string
    Timestamp    time.Time
}  

var consolidatedData = make(map[int]map[time.Time]UserHealthData)

func generateCompositeKey(userID int, timestamp time.Time) string {  
    return fmt.Sprintf("%d_%s", userID, timestamp.Format("20060102150405"))
}  

func addHealthData(userID int, heartRate int, activityLevel string, timestamp time.Time) {  
    //Initialize the inner map for the userID if it doesn't exist
    if _, ok := consolidatedData[userID]; !ok {
        consolidatedData[userID] = make(map[time.Time]UserHealthData)
    }
    _ := generateCompositeKey(userID, timestamp)
    consolidatedData[userID][timestamp] = UserHealthData{  
        HeartRate:    heartRate,
        ActivityLevel: activityLevel,
        Timestamp:    timestamp,
    }  
}  
func cleanupStaleData(retentionDuration time.Duration) {  
    now := time.Now()
    for userID, dataMap := range consolidatedData {  
        for timestamp, data := range dataMap {
            if now.Sub(data.Timestamp) > retentionDuration {  
                delete(consolidatedData[userID], timestamp)
            }
        }
        // Delete the user ID from the outer map if it has no remaining entries
        if len(consolidatedData[userID]) == 0 {
            delete(consolidatedData, userID)
        }
    }  
}  
func getHistoricalEntries(userID int) ([]UserHealthData, error) {  
    var entries []UserHealthData
    if dataMap, ok := consolidatedData[userID]; ok {
        for _, data := range dataMap {
            entries = append(entries, data)
        }  
        return entries, nil
    }  
    return nil, fmt.Errorf("user ID %d not found", userID)
}


func main() {  
    // Example data
    addHealthData(1, 72, "Active", time.Now().Add(-48*time.Hour))
    addHealthData(1, 75, "Moderately Active", time.Now().Add(-24*time.Hour))
    addHealthData(2, 65, "Sedentary", time.Now().Add(-72*time.Hour))
    addHealthData(3, 80, "Active", time.Now())
    
    // Query all entries for user ID 1
    if entries, err := getHistoricalEntries(1); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Historical entries for User ID 1:")
        for _, entry := range entries {  
            fmt.Printf("Data: %+v\n", entry)
        }
    }  
    cleanupStaleData(48 * time.Hour)
   
}  