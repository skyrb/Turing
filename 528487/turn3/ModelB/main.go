package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

// Our secret key for encryption (in real world use, use a secure key management system)
var encryptionKey = []byte("your-32-byte-encryption-key")

// HealthData holds the health metrics for a patient at a specific time.
type HealthData struct {
	EncryptedTimestamp []byte
	EncryptedWeight     []byte
	EncryptedBloodPressure []byte
	EncryptedGlucoseLevel []byte
}

// PatientRecord holds the health records for a patient.
type PatientRecord struct {
	ID     string
	Records map[string]HealthData
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

// EncryptData encrypts the plaintext data using AES-256.
func encryptData(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// DecryptData decrypts the encrypted data using AES-256.
func decryptData(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// AddRecord adds a new health record for a patient.
func (s *HealthService) AddRecord(patientID string, data HealthData) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	record, exists := s.records[patientID]
	if !exists {
		record = PatientRecord{
			ID:      patientID,
			Records: make(map[string]HealthData),
		}
	}
	record.Records[string(data.EncryptedTimestamp)] = data