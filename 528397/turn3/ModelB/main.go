package main

import (
	"fmt"
	"time"
)

const (
	maxRetryAttempts = 3 // Maximum number of retry attempts for each image
	retryDelay       = 5 * time.Second  // Delay between retry attempts
	circuitBreakTime  = 30 * time.Second // Duration for which circuit breaker is open
)

// Circuit breaker state
type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	Open
	HalfOpen
)

// Circuit breaker struct
type CircuitBreaker struct {
	state         CircuitBreakerState
	failureCount  int
	lastFailure   time.Time
	mutex         sync.Mutex
}

func (cb *CircuitBreaker) isOpen() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	return cb.state == Open
}

func (cb *CircuitBreaker) trip() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.state = Open
	cb.failureCount = 0
	cb.lastFailure = time.Now()
}

func (cb *CircuitBreaker) reset() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.state = Closed
	cb.failureCount = 0
}

func (cb *CircuitBreaker) shouldRetry() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == Open {
		// If the circuit breaker is open, check if the failure time has passed
		if time.Since(cb.lastFailure) >= circuitBreakTime {
			cb.state = HalfOpen
			cb.failureCount = 0
		} else {
			return false
		}
	}
	return true
}

func (cb *CircuitBreaker) onSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.failureCount = 0
}

func (cb *CircuitBreaker) onFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.failureCount++
	if cb.failureCount >= maxRetryAttempts {
		cb.trip()
	}
}

// Function to simulate distribution of product images with retry and circuit breaker mechanism
func distributeProductImages(region string, cb *CircuitBreaker) {
	for i := 1; ; i++ {
		fmt.Printf("Distributing product images to %s region (attempt %d)...\n", region, i)

		// Simulate image update to the primary CDN (London CDN)
		time.Sleep(2 * time.Second)

		// If the update is successful, break the loop and reset the circuit breaker
		if err := simulateImageUpdate("London"); err == nil {
			fmt.Printf("Product images distributed successfully to %s region\n", region)
			cb.onSuccess()
			return
		}

		// Check if the circuit breaker is open, if yes, go to the next attempt
		if !cb.shouldRetry() {
			fmt.Printf("Circuit breaker is open for %s region, skipping retry.\n", region)
			return
		}

		// If the update fails, wait for the retry delay and then try the secondary CDN (European CDN)
		fmt.Printf("Product image update to %s region failed, retrying...\n", region)
		cb.onFailure()
		time.Sleep(retryDelay)
		
		// Simulate image update to the secondary CDN
		if err := simulateImageUpdate("European"); err == nil {
			fmt.Printf("Product images distributed successfully to secondary European CDN for %s region\n", region)
			cb.onSuccess()
			return
		}
	}
}

// Simulate image update function (placeholder for actual CDN update logic)
func simulateImageUpdate(cdn string) error {
	// Introduce random failure to simulate image update failures
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) == 0 {
		return fmt.Errorf("image update failed to %s CDN", cdn)
	}
	return nil
} 
 