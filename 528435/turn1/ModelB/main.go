package main  
import (  
    "fmt"
    "log"
    "sync"
    "time"
)
// Device represents a device in the smart home.
type Device struct {  
    Name string
    Initialized bool
}
// Init initializes the device and sets its status to true.
func (d *Device) Init() {  
    time.Sleep(2 * time.Second) // Simulate device initialization
    d.Initialized = true
}
// InitDevices initializes all devices and waits for them to complete.
func InitDevices(devices []Device) {  
    var wg sync.WaitGroup
    wg.Add(len(devices))
    for i := range devices {  
        go func(i int) {  
            defer wg.Done()
            devices[i].Init()
            log.Printf("%s initialization completed.", devices[i].Name)
        }(i)
    }
    wg.Wait()
    log.Println("All devices are initialized.")
}
// LogDeviceStatus logs the status of each device.
func LogDeviceStatus(devices []Device) {  
    log.Println("Device status:")
    for _, d := range devices {  
        log.Printf("%s: Initialized = %t", d.Name, d.Initialized)
    }
}
func main() {  
    devices := []Device{
        {Name: "Light", Initialized: false},
        {Name: "Thermostat", Initialized: false},  
        {Name: "Security Camera", Initialized: false},
    }
    InitDevices(devices)
    LogDeviceStatus(devices)
    fmt.Println("Smart home system is fully operational.")
}  