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
	FirmwareUpdate() error
}

// Light device
type Light struct {
	id string
}

func (l *Light) Setup(dependenciesDone <-chan struct{}) error {
	time.Sleep(2 * time.Second)
	log.Printf("Light %s setup complete", l.id)
	return nil
}

func (l *Light) Name() string {
	return fmt.Sprintf("Light %s", l.id)
}

func (l *Light) FirmwareUpdate() error {
	// Simulate firmware update time
	time.Sleep(1 * time.Second)
	log.Printf("Light %s firmware update complete", l.id)
	return nil
}

// Camera device
type Camera struct {
	id string
}

func (c *Camera) Setup(dependenciesDone <-chan struct{}) error {
	<-dependenciesDone
	time.Sleep(1 * time.Second)
	log.Printf("Camera %s setup complete", c.id)
	return nil
}

func (c *Camera) Name() string {
	return fmt.Sprintf("Camera %s", c.id)
}

func (c *Camera) FirmwareUpdate() error {
	time.Sleep(1 * time.Second)
	log.Printf("Camera %s firmware update complete", c.id)
	return nil
}

// Function to batch update firmware
func batchFirmwareUpdate(devices []Device, batchSize int) {
	var wg sync.WaitGroup
	deviceChan := make(chan Device, len(devices))

	// Load devices into the channel
	for _, device := range devices {
		deviceChan <- device
	}
	close(deviceChan)

	for {
		// Process one batch at a time
		batch := make([]Device, 0, batchSize)
		for i := 0; i < batchSize; i++ {
			device, ok := <-deviceChan
			if !ok {
				break
			}
			batch = append(batch, device)
		}

		// If no devices left to process, break the loop
		if len(batch) == 0 {
			break
		}

		// Update firmware for the current batch
		for _, device := range batch {
			wg.Add(1)
			go func(dev Device) {
				defer wg.Done()
				if err := dev.FirmwareUpdate(); err != nil {
					log.Printf("Error updating firmware for device %s: %v", dev.Name(), err)
				}
			}(device)
		}

		// Wait for the current batch to complete
		wg.Wait()
		log.Printf("Completed firmware update for a batch of %d devices", len(batch))
	}

	log.Println("All devices firmware update process is complete.")
}

func main() {
	devices := []Device{
		&Light{id: "1"},
		&Light{id: "2"},
		&Camera{id: "1"},
		&Camera{id: "2"},
	}

	// Initialize devices (Assume all setup tasks have dependency logic handled)
	setupDevices(devices)
	log.Println("Devices initialized. Starting firmware updates.")

	// Update firmware in batches of 2
	batchFirmwareUpdate(devices, 2)
}

func setupDevices(devices []Device) {
	var wg sync.WaitGroup
	lightsDone := make(chan struct{})

	go func() {
		var lightsWG sync.WaitGroup
		for _, dev := range devices {
			if _, ok := dev.(*Light); ok {
				lightsWG.Add(1)
				go func(device Device) {
					defer lightsWG.Done()
					device.Setup(nil)
				}(dev)
			}
		}
		lightsWG.Wait()
		close(lightsDone)
	}()

	for _, device := range devices {
		if _, ok := device.(*Light); !ok {
			wg.Add(1)
			go func(dev Device) {
				defer wg.Done()
				dev.Setup(lightsDone)
			}(device)
		}
	}

	wg.Wait()
}