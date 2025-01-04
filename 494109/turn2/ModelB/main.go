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

const (
	// lockFilePermission is the permission to set on the lock file.
	lockFilePermission os.FileMode = 0644
)

// FileLock represents a lock on a file.
type FileLock struct {
	fileName string
	lockPath string
	// lockFile is the open file handle for the lock file.
	lockFile *os.File
	// lockMutex is used to synchronize access to lockFile.
	lockMutex sync.Mutex
}

// NewFileLock creates a new file lock.
func NewFileLock(fileName string) *FileLock {
	lockPath := fmt.Sprintf("%s.lock", fileName)
	return &FileLock{
		fileName: fileName,
		lockPath: lockPath,
	}
}

// Acquire attempts to acquire a lock on the file.
func (fl *FileLock) Acquire(ctx context.Context) (bool, error) {
	fl.lockMutex.Lock()
	defer fl.lockMutex.Unlock()

	if fl.lockFile != nil {
		// Lock is already acquired.
		return false, nil
	}

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
		}

		// Open the lock file in exclusive mode.
		lockFile, err := os.OpenFile(fl.lockPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, lockFilePermission)
		if err == nil {
			// Lock successfully acquired.
			fl.lockFile = lockFile
			return true, nil
		}

		// Check if the lock file exists and is locked by another process.
		if os.IsExist(err) {
			// Sleep for a short period to avoid busy waiting.
			time.Sleep(100 * time.Microsecond)
			continue
		}

		// Unexpected error.
		return false, err
	}
}

// Release releases a lock on the file.
func (fl *FileLock) Release() error {
	fl.lockMutex.Lock()
	defer fl.lockMutex.Unlock()

	if fl.lockFile == nil {
		// Lock is not acquired.
		return nil
	}

	// Close the lock file and remove it from the file system.
	if err := fl.lockFile.Close(); err != nil {
		return err
	}
	if err := os.Remove(fl.lockPath); err != nil {
		return err
	}
	fl.lockFile = nil
	return nil
}

// IsLocked checks if the file is currently locked.
func (fl *FileLock) IsLocked() (bool, error) {
	fl.lockMutex.Lock()
	defer fl.lockMutex.Unlock()

	_, err := os.Stat(fl.lockPath)
	return err == nil, err
}

// worker performs some work while holding the lock.
func worker(lock *FileLock, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if acquired, err := lock.Acquire(ctx); err != nil {
		log.Printf("Error acquiring lock: %v", err)
		return
	} else if !acquired {
		log.Printf("Lock acquisition timed out")
		return
	}
	defer lock.Release()

	log.Printf("Lock acquired on file %s", lock.fileName)
