
package main  
import (  
    "fmt"
    "runtime"
    "sync"
    "time"
)
// Logger is a custom logger that can log any anomaly occurred during the map operation
type Logger struct {  
    mutex sync.Mutex  
    logEntries []LogEntry  
}  

// LogEntry is an entry in the logger, containing details of the anomaly
type LogEntry struct {  
    Anomaly string    `json:"anomaly"`  
    Goroutine string `json:"goroutine"`  
    Timestamp time.Time `json:"timestamp"`  
}  

// Log logs an anomaly along with the goroutine information and timestamp
func (l *Logger) Log(anomaly string) {  
    l.mutex.Lock()  
    defer l.mutex.Unlock()  
    pc, file, line, _ := runtime.Caller(1)  
    goroutine := fmt.Sprintf("%s:%d: %s", file, line, runtime.FuncForPC(pc).Name())  
    l.logEntries = append(l.logEntries, LogEntry{  
        Anomaly: anomaly,  
        Goroutine: goroutine,  
        Timestamp: time.Now(),  
    })  
}  

// AnomalyDetector is responsible for monitoring map operations and logging anomalies
type AnomalyDetector struct {  
    logger *Logger  
}  

// NewAnomalyDetector creates a new AnomalyDetector instance with a given logger
func NewAnomalyDetector(logger *Logger) *AnomalyDetector {  
    return &AnomalyDetector{logger: logger}  
}  

// wrapMap takes a regular map and returns a wrapped map that detects nil map access and logs it
func (d *AnomalyDetector) wrapMap(m map[int]int) map[int]int {  
    return map[int]int{  
        wrapKeyAccess(d.logger, nilKey):   wrapValueAccess(d.logger, nilMap),  
    }  
}  

// wrapKeyAccess wraps a key value to detect access to a nil key
func wrapKeyAccess(logger *Logger, originalKey int) int {  
    return func(originalKey)  {  
        if originalKey == nil {  
            logger.Log("Accessing nil key")  
        }  
        return originalKey  
    }()  
}  

// wrapValueAccess wraps a value value to detect access to a nil map
func wrapValueAccess(logger *Logger, originalValue int) int {  
    return func(originalValue) {  
        if originalValue == nil {  
            logger.Log("Accessing nil map")  
        }  
        return originalValue  
    }()  
}  

func main() {  
    logger := &Logger{}  
    anomalyDetector := NewAnomalyDetector(logger)  
    wrappedMap := anomalyDetector.wrapMap(map[int]int{1:10})  

    // Multiple goroutines try to access the shared wrapped map
    var wg sync.WaitGroup  
    wg.Add(2)  

    go func() {  
        defer wg.Done()  
        wrappedMap[1]++  
    }()  

    go func() {  
        defer wg.Done()  
        wrappedMap[2]++  
    }()  

    wg.Wait()  

    // After the goroutines have finished, print the log entries
    for _, entry := range logger.logEntries {  
        fmt.Println(entry)  
    }  
}
  