package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

type Device interface {
    Setup(wg *sync.WaitGroup) error
    Name() string
}

type Light struct {
    id string
}

func (l *Light) Setup(wg *sync.WaitGroup) error {
    defer wg.Done()
    time.Sleep(2 * time.Second)
    log.Printf("Light %s setup complete", l.id)
    return nil
}

func (l *Light) Name() string {
    return fmt.Sprintf("Light %s", l.id)
}

type Thermostat struct {
    id string
}

func (t *Thermostat) Setup(wg *sync.WaitGroup) error {
    defer wg.Done()
    time.Sleep(3 * time.Second)
    log.Printf("Thermostat %s setup complete", t.id)
    return nil
}

func (t *Thermostat) Name() string {
    return fmt.Sprintf("Thermostat %s", t.id)
}

type Camera struct {
    id     string
    light1 *Light
}

func (c *Camera) Setup(wg *sync.WaitGroup) error {
    defer wg.Done()
    wg.Add(1)
    go func() {
        defer wg.Done()
        <-time.After(2 * time.Second)
        log.Printf("Camera %s setup complete after waiting for lights", c.id)
    }()
    log.Printf("Camera %s setup initiated, waiting for lights...", c.id)
    return nil
}

func (c *Camera) Name() string {
    return fmt.Sprintf("Camera %s", c.id)
}

func main() {
    devices := []Device{
        &Light{id: "1"},
        &Thermostat{id: "1"},
        &Camera{id: "1", light1: &Light{id: "1"}},
    }

    var wg sync.WaitGroup
    statusChan := make(chan string, len(devices))

    for _, device := range devices {
        wg.Add(1)
        go func(dev Device) {
            defer wg.Done()
            if err := dev.Setup(&wg); err != nil {
                log.Printf("Error setting up device %s: %v", dev.Name(), err)
                return
            }
            statusChan <- fmt.Sprintf("%s setup completed successfully", dev.Name())
        }(device)
    }

    wg.Wait()
    close(statusChan)

    log.Println("All devices are set up. System is now fully operational")

    for status := range statusChan {
        log.Println(status)
    }
}
