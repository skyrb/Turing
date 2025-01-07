package main  
import ("fmt")

type Activity struct {
    Name       string
    Duration   int   // In minutes
    HeartRate  int
}

func main() {
    // Initialize an empty slice of Activity
    var activities []Activity

    // Append some activities to the slice
    activities = append(activities, 
        Activity{Name: "Running", Duration: 30, HeartRate: 140},
        Activity{Name: "Yoga", Duration: 45, HeartRate: 90},
        Activity{Name: "Cycling", Duration: 60, HeartRate: 160},
        Activity{Name: "Swimming", Duration: 25, HeartRate: 100},
        Activity{Name: "Walking", Duration: 35, HeartRate: 80},
    )

    // Filter and display all activities where heart rate exceeds 120
    fmt.Println("Activities with heart rate exceeding 120:")
    for _, activity := range activities {
        if activity.HeartRate > 120 {
            fmt.Printf("%s - Duration: %d minutes, Heart Rate: %d\n", activity.Name, activity.Duration, activity.HeartRate)
        }
    }
}