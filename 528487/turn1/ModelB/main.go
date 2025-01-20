
package main  
import (  
    "fmt"
    "time"
)
type PatientHealthData struct {  
    Weight float64
    BloodPressure float64
    GlucoseLevel float64
}
type HealthDataService struct {  
    patientData map[string]map[time.Time]PatientHealthData
}
func NewHealthDataService() *HealthDataService {  
    return &HealthDataService{
        patientData: make(map[string]map[time.Time]PatientHealthData),
    }
}
func (hds *HealthDataService) AddPatientHealthData(patientID string, data PatientHealthData) {  
    if hds.patientData[patientID] == nil {  
        hds.patientData[patientID] = make(map[time.Time]PatientHealthData)
    }
    now := time.Now()
    hds.patientData[patientID][now] = data
}
func (hds *HealthDataService) GetPatientHealthData(patientID string) (map[time.Time]PatientHealthData, bool) {  
    data, ok := hds.patientData[patientID]
    return data, ok
}
func (hds *HealthDataService) GetPatientHealthDataByDate(patientID string, date time.Time) (PatientHealthData, bool) {  
    data, ok := hds.patientData[patientID]
    if !ok {
        return PatientHealthData{}, false
    }
    healthData, ok := data[date]
    return healthData, ok
}
func main() {  
    hds := NewHealthDataService()
    hds.AddPatientHealthData("patient1", PatientHealthData{Weight: 70.5, BloodPressure: 120.8, GlucoseLevel: 90.0})
    hds.AddPatientHealthData("patient2", PatientHealthData{Weight: 65.0, BloodPressure: 115.4, GlucoseLevel: 85.0})
    hds.AddPatientHealthData("patient1", PatientHealthData{Weight: 71.0, BloodPressure: 122.0, GlucoseLevel: 92.0})
    patient1Data, ok := hds.GetPatientHealthData("patient1")
    if ok {
        fmt.Println("Patient 1 Health Data:")
        for date, data := range patient1Data {
            fmt.Printf("Date: %s, Weight: %.1f, Blood Pressure: %.1f, Glucose Level: %.1f\n", date.Format("2006-01-02"), data.Weight, data.BloodPressure, data.GlucoseLevel)
        }
    } else {
        fmt.Println("Patient 1 not found.")
    }
    patient2Data, ok := hds.GetPatientHealthData("patient2")  
    if ok {
        fmt.Println("\nPatient 2 Health Data:")
        for date, data := range patient2Data {
            fmt.Printf("Date: %s, Weight: %.1f, Blood Pressure: %.1f, Glucose Level: %.1f\n", date.Format("2006-01-02"), data.Weight, data.BloodPressure, data.GlucoseLevel)
        }
    } else {
        fmt.Println("Patient 2 not found.")
    }
    date := time.Date(2023, 07, 31, 0, 0, 0, 0, time.UTC)
    patient1DataByDate, ok := hds.GetPatientHealthDataByDate("patient1", date)
    if ok {
        fmt.Println("\nPatient 1 Health Data on", date.Format("2006-01-02"))
        fmt.Printf("Weight: %.1f, Blood Pressure: %.1f, Glucose Level: %.1f\n", patient1DataByDate.Weight, patient1DataByDate.BloodPressure, patient1DataByDate.GlucoseLevel)
    } else {
        fmt.Println("Patient 1 data not found for", date.Format("2006-01-02"))
    }
}
  