package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

// FileLock represents a lock on a file
type FileLock struct {
	fileName string
	mutex    sync.Mutex
	lockPath string
}

// NewFileLock creates a new file lock
func NewFileLock(fileName string) *FileLock {
	lockPath := fmt.Sprintf("%s.lock", fileName)
	return &FileLock{
		fileName: fileName,
		lockPath: lockPath,
	}
}

// Acquire attempts to acquire a lock on the file
func (fl *FileLock) Acquire(ctx context.Context) (bool, error) {
	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// Check if the lock file already exists
	if _, err := os.Stat(fl.lockPath); err == nil {
		return false, fmt.Errorf("lock already exists on file %s", fl.fileName)
	}

	// Create a new lock file
	if err := ioutil.WriteFile(fl.lockPath, []byte("locked"), 0644); err != nil {
		return false, err
	}

	return true, nil
}

// Release releases a lock on the file
func (fl *FileLock) Release() error {
	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	// Check if the lock file exists
	if _, err := os.Stat(fl.lockPath); os.IsNotExist(err) {
		return fmt.Errorf("lock file does not exist for file %s", fl.fileName)
	}

	// Remove the lock file
	if err := os.Remove(fl.lockPath); err != nil {
		return err
	}

	return nil
}

// IsLocked checks if the file is currently locked
func (fl *FileLock) IsLocked() (bool, error) {
	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	// Check if the lock file exists
	if _, err := os.Stat(fl.lockPath); err == nil {
		return true, nil
	}

	return false, nil
}

func main() {
	fileName := "example.txt"
	lock := NewFileLock(fileName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if acquired, err := lock.Acquire(ctx); err != nil {
		log.Fatalf("Error acquiring lock: %v", err)
	} else if !acquired {
		log.Fatalf("Lock acquisition timed out")
	}

	log.Printf("Lock acquired on file %s", fileName)

	time.Sleep(2 * time.Second) // Simulate some work

	if err := lock.Release(); err != nil {
		log.Fatalf("Error releasing lock: %v", err)
	}

	log.Printf("Lock released on file %s", fileName)
}