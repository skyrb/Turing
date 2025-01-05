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
	lockPath string
	sync.Mutex
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
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if _, err := os.Stat(fl.lockPath); err == nil {
		return false, fmt.Errorf("lock already exists on file %s", fl.fileName)
	}

	// Use a mutex to ensure thread-safe lock creation
	fl.Lock()
	defer fl.Unlock()

	if _, err := os.Stat(fl.lockPath); err == nil {
		return false, fmt.Errorf("lock already exists on file %s", fl.fileName)
	}

	if err := ioutil.WriteFile(fl.lockPath, []byte("locked"), 0644); err != nil {
		return false, err
	}

	return true, nil
}

// Release releases a lock on the file
func (fl *FileLock) Release() error {
	if _, err := os.Stat(fl.lockPath); os.IsNotExist(err) {
		return fmt.Errorf("lock file does not exist for file %s", fl.fileName)
	}

	// Use a mutex to ensure thread-safe lock removal
	fl.Lock()
	defer fl.Unlock()

	if _, err := os.Stat(fl.lockPath); err == nil {
		return fmt.Errorf("lock already exists on file %s", fl.fileName)
	}

	if err := os.Remove(fl.lockPath); err != nil {
		return err
	}

	return nil
}

// IsLocked checks if the file is currently locked
func (fl *FileLock) IsLocked() (bool, error) {
	if _, err := os.Stat(fl.lockPath); err == nil {
		return true, nil
	}

	return false, nil
}

func main() {
	fileName := "example.txt"
	lock := NewFileLock(fileName)

	// Simulate concurrent access
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ { // Simulate 100 concurrent workers
		wg.Add(1)
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if acquired, err := lock.Acquire(ctx); err != nil {
				log.Fatalf("Error acquiring lock: %v", err)
			} else if !acquired {
				log.Fatalf("Lock acquisition timed out")
			}

			log.Printf("Lock acquired on file %s", fileName)

			time.Sleep(1 * time.Second) // Simulate some work

			if err := lock.Release(); err != nil {
				log.Fatalf("Error releasing lock: %v", err)
			}

			log.Printf("Lock released on file %s", fileName)

			wg.Done()
		}()
	}

	wg.Wait()
	log.Printf("All locks acquired and released successfully.")
}