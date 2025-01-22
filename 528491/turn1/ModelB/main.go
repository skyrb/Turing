
package main  
import (  
    "fmt"
    "log"
    "strings"
    "sync"

    "github.com/eclipse/paho.mqtt.golang"
)

// WeatherData stores climate data
type WeatherData struct {
    Temperature float32 `json:"temperature"`
    Humidity    float32 `json:"humidity"`
}

var (  
    // Map to store climate data keyed by region name
    climateDataMap = make(map[string]WeatherData)
    dataMapLock    sync.Mutex
)

func main() {  
    regions := []string{"region1", "region2", "region3"}
    // Subscribe to MQTT topics for each region
    for _, region := range regions {
        go subscribeToRegion(region)
    }

    // Block the main routine from exiting to keep the subscription running
    select {}
}

func subscribeToRegion(region string) {  
    client := createMQTTClient(region)
    topic := fmt.Sprintf("weather/%s", region)

    // Subscribe to the MQTT topic
    if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
        log.Fatalf("Error subscribing to topic %s: %s", topic, token.Error())
    }

    fmt.Printf("Subscribed to region %s, listening for data...\n", region)
}

func createMQTTClient(region string) mqtt.Client {  
    brokerURL := fmt.Sprintf("tcp://mqtt.example.com:1883")
    clientID := fmt.Sprintf("weather-collector-%s", region)

    connOpts := mqtt.NewClientOptions().AddBroker(brokerURL).SetClientID(clientID)
    connOpts.SetDefaultPublishHandler(messageHandler)
    connOpts.SetConnectionLostHandler(connLostHandler)

    client := mqtt.NewClient(connOpts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Error connecting to MQTT broker: %s", token.Error())
    }

    return client
}

func messageHandler(client mqtt.Client, message mqtt.Message) {  
    // Parse the topic to get the region name
    topic := message.Topic()
    region := strings.TrimPrefix(topic, "weather/")

    // Unmarshal the JSON message into a WeatherData struct
    var data WeatherData
   // Using json library to unmarshal
    // _ := json.Unmarshal(message.Payload(), &data)
    log.Printf("Received message from region %s: %s\n", region, message.Payload())

    // Update the climate data map with the new data
    dataMapLock.Lock()
    climateDataMap[region] = data
    dataMapLock.Unlock()
}

func connLostHandler(client mqtt.Client, err error) {  
    log.Printf("Connection lost: %s\n", err.Error())
}
