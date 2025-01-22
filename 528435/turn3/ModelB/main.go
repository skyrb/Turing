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
	UpdateFirmware(batchDone <-chan struct{}) error
	Name() string
}

// Light device
type Light struct {
	id     string
	updated bool
}

func (l *Light) Setup(dependenciesDone <-chan struct{}) error {
	// Simulate setup time
	time.Sleep(2 * time.Second)
	log.Printf("Light %s setup complete", l.id)
	return nil
}

func (l *Light) UpdateFirmware(batchDone <-chan struct{}) error {
	// Wait for previous batch to complete
	<-batchDone

	if !l.updated {
		// Simulate firmware update time
		time.Sleep(1 * time.Second)
		log.Printf("Light %s firmware update completed", l.id)
		l.updated = true
	}
	return nil
}

func (l *Light) Name() string {
	return fmt.Sprintf("Light %s", l.id)
}

// Similar implementations for Thermostat and Camera devices

func main() {
	devices := []Device{
		&Light{id: "1"},
		&Light{id: "2"},
		&Thermostat{id: "1"},
		&Camera{id: "1"},
	}

	var wg sync.WaitGroup
	statusChan := make(chan string, len(devices))
	lightsDone := make(chan struct{}) // Signal channel for light setup completion

	// Setup lights and signal completion
	go func() {
		var lightsWG sync.WaitGroup
		for _, device := range devices {
			if _, isLight := device.(*Light); isLight {
				lightsWG.Add(1)
				go func(dev Device) {
					defer lightsWG.Done()
					if err := dev.Setup(nil); err != nil {
						log.Printf("Error setting up device %s: %v", dev.Name(), err)
						return
					}
					statusChan <- fmt.Sprintf("%s setup completed successfully", dev.Name())
				}(device)
			}
		}
		lightsWG.Wait()
		close(lightsDone) // Signal that all lights are set up
	}()

	// Function to update firmware in batches
	updateFirmwareInBatches := func() {
		batchDone := make(chan struct{})
		defer close(batchDone)

		for _, device := range devices {
			wg.Add(1)
			go func(dev Device) {
				defer wg.Done()
				if err := dev.UpdateFirmware(batchDone); err != nil {
					log.Printf("Error updating firmware for device %s: %v", dev.Name(), err)
					return
				}
				statusChan <- fmt.Sprintf("%s firmware updated successfully", dev.Name())
			}(device)
		}
	}

	// Start initial setup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, device := range devices {
			if err := device.Setup(lightsDone); err != nil {
				log.Printf("Error setting up device %s: %v", device.Name(), err)
				return
			}
			statusChan <- fmt.Sprintf("%s setup completed successfully", device.Name())
		}
	}()
	wg.Wait()

	// Log the statuses
	for status := range statusChan {
		log.Println(status)
	}

	// Periodically update firmware in batches
	for {
		updateFirmwareInBatches()
		time.Sleep(time.Duration(30) * time.Second) // Update every 30 seconds
	}
}