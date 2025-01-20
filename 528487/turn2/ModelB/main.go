package main  

import (  
    "fmt"
    "sort"
    "sync"
    "time"
)

// HealthData holds the health metrics for a patient at a specific time.
type HealthData struct {
	Timestamp     time.Time
	Weight        float64
	BloodPressure string
	GlucoseLevel  float64
}

// PatientRecord holds the health records for a patient.
type PatientRecord struct {
	ID     string
	Records map[time.Time]HealthData
}

// HealthService manages patient records.
type HealthService struct {
	mutex   sync.RWMutex
	records map[string]PatientRecord
}

// NewHealthService creates a new HealthService.
func NewHealthService() *HealthService {
	return &HealthService{
		records: make(map[string]PatientRecord),
	}
}

// AddRecord adds a new health record for a patient.
func (s *HealthService) AddRecord(patientID string, data HealthData) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	record, exists := s.records[patientID]
	if !exists {
		record = PatientRecord{
			ID:      patientID,
			Records: make(map[time.Time]HealthData),
		}
	}
	record.Records[data.Timestamp] = data
	s.records[patientID] = record
}

// GetRecords retrieves all health records for a patient.
func (s *HealthService) GetRecords(patientID string) ([]HealthData, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	record, exists := s.records[patientID]
	if !exists {
		return nil, false
	}

	var records []HealthData
	for _, data := range record.Records {
		records = append(records, data)
	}
	return records, true
}

// UpdateRecord updates an existing health record for a patient.
func (s *HealthService) UpdateRecord(patientID string, timestamp time.Time, data HealthData) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	record, exists := s.records[patientID]
	if !exists {
		return false
	}

	if _, exists := record.Records[timestamp]; !exists {
		return false
	}

	record.Records[timestamp] = data
	s.records[patientID] = record
	return true
}

// DeleteRecord deletes a health record for a patient.
func (s *HealthService) DeleteRecord(patientID string, timestamp time.Time) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	record, exists := s.records[patientID]
	if !exists {
		return false
	}

	if _, exists := record.Records[timestamp]; !exists {
		return false
	}
	delete(record.Records, timestamp)
	return true
}

func (s *HealthService) QueryBloodPressureTrend(patientID string, lastMonths int) (bool, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    record, found := s.records[patientID]
    if !found {
        return false, fmt.Errorf("patient ID not found: %s", patientID)
    }
    // Calculate the time range for the last months
    endTime := time.Now()
    startTime := endTime.AddDate(0, -lastMonths, 0)

    // Filter the records within the specified time range
    var filteredRecords []HealthData
    for timestamp, data := range record.Records {
        if timestamp.After(startTime) && timestamp.Before(endTime) {
            filteredRecords = append(filteredRecords, data)
        }
    }
    // Check if the number of records is enough for analysis
    if len(filteredRecords) < 3 {
        return false, fmt.Errorf("not enough records for trend analysis")
    }
    // Sort the records by timestamp in ascending order
    sort.Slice(filteredRecords, func(i, j int) bool {
        return filteredRecords[i].Timestamp.Before(filteredRecords[j].Timestamp)
    })
    // Analyze the trend of blood pressure
    for i := 1; i < len(filteredRecords); i++ {
        currentBP, err := parseBloodPressure(filteredRecords[i].BloodPressure)
        if err != nil {
            return false, err
        }
        prevBP, err := parseBloodPressure(filteredRecords[i-1].BloodPressure)
        if err != nil {
            return false, err
        }
        if currentBP.Systolic <= prevBP.Systolic {
            return false, nil // Not increasing consistently
        }
    }
    return true, nil // Increasing consistently
}

// parseBloodPressure extracts systolic and diastolic values from the blood pressure string
func parseBloodPressure(bp string) (systolic, diastolic int, err error) {
    fmt.Sscanf(bp, "%d/%d", &systolic, &diastolic)
    return
}

func main() {
	service := NewHealthService()

	// Example usage
	timestamp := time.Now()
	data := HealthData{
		Timestamp:     timestamp,
		Weight:        70.5,
		BloodPressure: "120/80",
		GlucoseLevel:  5.5,
	}

	// Add a record
	service.AddRecord("patient1", data)

	// Retrieve records
	records, found := service.GetRecords("patient1")
	if found {
		fmt.Println("Records for patient1:")
		for _, record := range records {
			fmt.Printf("%v\n", record)
		}
	} else {
		fmt.Println("No records found for patient1.")
	}

    // Example usage for trend analysis:
    consistentlyIncreasing, err := service.QueryBloodPressureTrend("patient1", 3)
    if err != nil {
        fmt.Printf("Error analyzing trend: %v\n", err)
    } else if consistentlyIncreasing {
        fmt.Println("Patient's blood pressure is consistently increasing.")
    } else {
        fmt.Println("Patient's blood pressure trend is not consistent.")
    }
}