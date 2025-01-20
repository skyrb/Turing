package main

import (
	"fmt"
	"sync"
	"time"
)

// HealthData holds the health metrics for a patient at a specific time.
type HealthData struct {
	Timestamp     time.Time
	Weight        float64
	BloodPressure int // In this example, we use an integer to reflect systolic pressure for simplicity.
	GlucoseLevel  float64
}

// PatientRecord holds the health records for a patient.
type PatientRecord struct {
	ID      string
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

// IdentifyIncreasingBP detects patients with consistently increasing blood pressure over the last three months.
func (s *HealthService) IdentifyIncreasingBP() []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var patientsWithIncreasingBP []string
	threeMonthsAgo := time.Now().AddDate(0, -3, 0)

	for id, record := range s.records {
		if s.isBloodPressureIncreasing(record.Records, threeMonthsAgo) {
			patientsWithIncreasingBP = append(patientsWithIncreasingBP, id)
		}
	}

	return patientsWithIncreasingBP
}

// isBloodPressureIncreasing checks if the patient's blood pressure has been consistently increasing.
func (s *HealthService) isBloodPressureIncreasing(records map[time.Time]HealthData, cutoff time.Time) bool {
	var sortedRecords []HealthData

	// Filter and sort records by timestamp
	for _, data := range records {
		if data.Timestamp.After(cutoff) {
			sortedRecords = append(sortedRecords, data)
		}
	}

	if len(sortedRecords) < 2 {
		return false
	}

	// Ensure sorted by time
	sort.Slice(sortedRecords, func(i, j int) bool {
		return sortedRecords[i].Timestamp.Before(sortedRecords[j].Timestamp)
	})

	// Check for consistent increase
	for i := 1; i < len(sortedRecords); i++ {
		if sortedRecords[i].BloodPressure <= sortedRecords[i-1].BloodPressure {
			return false
		}
	}

	return true
}

func main() {
	service := NewHealthService()

	// Example records setup
	timestamps := []time.Time{
		time.Now().AddDate(0, -3, 0),
		time.Now().AddDate(0, -2, 0),
		time.Now().AddDate(0, -1, 0),
	}
	for _, ts := range timestamps {
		data := HealthData{
			Timestamp:     ts,
			Weight:        70.5,
			BloodPressure: 120 + ts.Month(), // Increase mock systolic values for illustrative purposes.
			GlucoseLevel:  5.5,
		}
		service.AddRecord("patient1", data)
	}

	increasingBPPatients := service.IdentifyIncreasingBP()
	fmt.Println("Patients with consistently increasing blood pressure over the last three months:", increasingBPPatients)
}