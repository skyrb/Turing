package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Device interface representing a smart device in the home
type Device interface {
	Setup(dependenciesDone <-chan struct{}) error
	Name() string
}

// Light device
type Light struct {
	id string
}

func (l *Light) Setup(dependenciesDone <-chan struct{}) error {
	// Simulate setup time
	time.Sleep(2 * time.Second)
	log.Printf("Light %s setup complete", l.id)
	return nil
}

func (l *Light) Name() string {
	return fmt.Sprintf("Light %s", l.id)
}

// Thermostat device
type Thermostat struct {
	id string
}

func (t *Thermostat) Setup(dependenciesDone <-chan struct{}) error {
	// Simulate setup time
	time.Sleep(3 * time.Second)
	log.Printf("Thermostat %s setup complete", t.id)
	return nil
}

func (t *Thermostat) Name() string {
	return fmt.Sprintf("Thermostat %s", t.id)
}

// Camera device
type Camera struct {
	id string
}

func (c *Camera) Setup(dependenciesDone <-chan struct{}) error {
	// Wait for lights to complete setup
	<-dependenciesDone

	// Simulate setup time
	time.Sleep(1 * time.Second)
	log.Printf("Camera %s setup complete", c.id)
	return nil
}

func (c *Camera) Name() string {
	return fmt.Sprintf("Camera %s", c.id)
}

func main() {
	devices := []Device{
		&Light{id: "1"},
		&Light{id: "2"},
		&Thermostat{id: "1"},
		&Camera{id: "1"},
	}

	var wg sync.WaitGroup
	statusChan := make(chan string, len(devices))
	lightsDone := make(chan struct{})

	// Start the setup process for each device concurrently
	for _, device := range devices {
		wg.Add(1)
		go func(dev Device) {
			defer wg.Done()
			// Lights signal completion through the lightsDone channel
			if _, isLight := dev.(*Light); isLight {
				if err := dev.Setup(nil); err != nil {
					log.Printf("Error setting up device %s: %v", dev.Name(), err)
					return
				}
				statusChan <- fmt.Sprintf("%s setup completed successfully", dev.Name())
				lightsDone <- struct{}{}
			} else {
				if err := dev.Setup(lightsDone); err != nil {
					log.Printf("Error setting up device %s: %v", dev.Name(), err)
					return
				}
				statusChan <- fmt.Sprintf("%s setup completed successfully", dev.Name())
			}
		}(device)
	}

	// Wait for all devices to finish setup
	wg.Wait()
	close(statusChan)
	close(lightsDone)

	log.Println("All devices are set up. System is now fully operational")

	// Log the statuses
	for status := range statusChan {
		log.Println(status)
	}
}