package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated function to get predicted priority of regions
func getPredictedContents() map[string]int {
	return map[string]int{
		"North America": 1, // High-demand due to Black Friday
		"India":         2, // High-demand due to Diwali
		"Europe":        3, // Medium-demand
		"Other":         4, // Low-demand
	}
}

type ContentTask struct {
	region   string
	priority int
}

var (
	highPriorityTasks   = make(chan ContentTask, 10)
	lowPriorityTasks    = make(chan ContentTask, 10)
	realTimeUpdateTasks = make(chan ContentTask, 10)
)

func preloadContent(task ContentTask, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Preloading content for %s with priority %d\n", task.region, task.priority)
	// Simulate processing time
	time.Sleep(time.Duration(2-task.priority) * time.Second)
	fmt.Printf("Completed preloading content for %s\n", task.region)
}

func handleRealTimeUpdates(task ContentTask, done chan bool) {
	fmt.Printf("Real-time update for %s under priority %d\n", task.region, task.priority)
	// Simulate rapid processing time
	time.Sleep(1 * time.Second)
	fmt.Printf("Completed real-time update for %s\n", task.region)
	done <- true
}

func main() {
	var wg sync.WaitGroup
	predictedContents := getPredictedContents()

	// Dispatch tasks based on priority
	for region, priority := range predictedContents {
		task := ContentTask{region: region, priority: priority}
		if priority <= 2 {
			wg.Add(1)
			highPriorityTasks <- task
		} else {
			wg.Add(1)
			lowPriorityTasks <- task
		}
	}

	// Simulate periodic real-time updates
	go func() {
		for {
			for region, priority := range predictedContents {
				realTimeUpdateTasks <- ContentTask{region: region, priority: priority}
				time.Sleep(5 * time.Second)
			}
		}
	}()

	// Worker to process high priority tasks
	go func() {
		for task := range highPriorityTasks {
			go preloadContent(task, &wg)
		}
	}()

	// Worker to process low priority tasks
	go func() {
		for task := range lowPriorityTasks {
			go preloadContent(task, &wg)
		}
	}()

	// Handle real-time updates non-blocking
	go func() {
		done := make(chan bool)
		for task := range realTimeUpdateTasks {
			go func(t ContentTask) {
				handleRealTimeUpdates(t, done)
			}(task)

			<-done // Wait for completion to balance the writes to realTimeUpdateTasks
		}
	}()

	wg.Wait()
	close(highPriorityTasks)
	close(lowPriorityTasks)
	close(realTimeUpdateTasks)

	fmt.Println("All content tasks and updates processed.")
}