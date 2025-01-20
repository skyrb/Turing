package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sort"
	"sync"
	"time"
)

type HealthData struct {
	Timestamp     time.Time
	Weight        float64
	EncryptedBP   string // Encrypted Blood Pressure
	EncryptedGlucose string // Encrypted Glucose Level
}

type PatientRecord struct {
	ID      string
	Records map[time.Time]HealthData
}

type HealthService struct {
	mutex   sync.RWMutex
	records map[string]PatientRecord
	key     []byte
}

func NewHealthService(encryptionKey []byte) *HealthService {
	return &HealthService{
		records: make(map[string]PatientRecord),
		key:     encryptionKey,
	}
}

func (s *HealthService) AddRecord(patientID string, data HealthData, bloodPressure int, glucoseLevel float64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	record, exists := s.records[patientID]
	if !exists {
		record = PatientRecord{
			ID:      patientID,
			Records: make(map[time.Time]HealthData),
		}
	}

	encryptedBP, err := encrypt(s.key, fmt.Sprintf("%d", bloodPressure))
	if err != nil {
		return err
	}

	encryptedGlucose, err := encrypt(s.key, fmt.Sprintf("%.2f", glucoseLevel))
	if err != nil {
		return err
	}

	data.EncryptedBP = encryptedBP
	data.EncryptedGlucose = encryptedGlucose
	record.Records[data.Timestamp] = data
	s.records[patientID] = record

	return nil
}

func encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return fmt.Sprintf("%x", ciphertext), nil
}

func decrypt(key []byte, cryptoText string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

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

func (s *HealthService) isBloodPressureIncreasing(records map[time.Time]HealthData, cutoff time.Time) bool {
	var sortedRecords []HealthData

	for _, data := range records {
		if data.Timestamp.After(cutoff) {
			sortedRecords = append(sortedRecords, data)
		}
	}

	if len(sortedRecords) < 2 {
		return false
	}

	sort.Slice(sortedRecords, func(i, j int) bool {
		return sortedRecords[i].Timestamp.Before(sortedRecords[j].Timestamp)
	})

	previousBP := -1
	for _, record := range sortedRecords {
		decryptedBP, err := decrypt(s.key, record.EncryptedBP)
		if err != nil {
			continue
		}

		bp, err := strconv.Atoi(decryptedBP)
		if err != nil {
			continue
		}

		if previousBP != -1 && bp <= previousBP {
			return false
		}
		previousBP = bp
	}

	return true
}

func main() {
	key := []byte("example key 1234") // Ensure key length is 16, 24, or 32 bytes for AES
	service := NewHealthService(key)

	timestamps := []time.Time{
		time.Now().AddDate(0, -3, 0),
		time.Now().AddDate(0, -2, 0),
		time.Now().AddDate(0, -1, 0),
	}
	for i, ts := range timestamps {
		err := service.AddRecord("patient1", HealthData{Timestamp: ts}, 120+i, 5.5+i)
		if err != nil {
			fmt.Println("Error adding record:", err)
		}
	}

	increasingBPPatients := service.IdentifyIncreasingBP()
	fmt.Println("Patients with consistently increasing blood pressure over the last three months:", increasingBPPatients)
}