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

	// Update a record
	updateData := HealthData{
		Timestamp:     timestamp,
		Weight:        71.0,
		BloodPressure: "118/79",
		GlucoseLevel:  5.2,
	}
	if service.UpdateRecord("patient1", timestamp, updateData) {
		fmt.Println("Record updated successfully.")
	} else {
		fmt.Println("Failed to update record.")
	}

	// Delete a record
	if service.DeleteRecord("patient1", timestamp) {
		fmt.Println("Record deleted successfully.")
	} else {
		fmt.Println("Failed to delete record.")
	}
}