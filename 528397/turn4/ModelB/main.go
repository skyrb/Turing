package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (  
    numRegions        = 3
    maxPreloadWorkers = 5
    preloadDelay      = 10 * time.Second
)

// Content represents a piece of content to be preloaded
type Content struct {  
    ID       string
    Region   string
    Priority int
}

// PriorityQueue manages content preloading tasks with priority
type PriorityQueue struct {  
    tasks []Content
}

func (pq *PriorityQueue) Len() int {  
    return len(pq.tasks)
}

func (pq *PriorityQueue) Less(i, j int) bool {  
    return pq.tasks[i].Priority > pq.tasks[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {  
    pq.tasks[i], pq.tasks[j] = pq.tasks[j], pq.tasks[i]
}

func (pq *PriorityQueue) Push(x interface{}) {  
    task := x.(Content)
    pq.tasks = append(pq.tasks, task)
}

func (pq *PriorityQueue) Pop() interface{} {  
    n := len(pq.tasks)
    task := pq.tasks[n-1]
    pq.tasks = pq.tasks[:n-1]
    return task
}

// PreloadContent simulates preloading content to a region
func PreloadContent(content Content) {  
    fmt.Printf("Preloading content %s to %s with priority %d\n", content.ID, content.Region, content.Priority)
    time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func preloadWorker(pq *PriorityQueue, wg *sync.WaitGroup) {  
    defer wg.Done()
    for {  
        if pq.Len() == 0 {  
            time.Sleep(preloadDelay)
            continue
        }
        content := heap.Pop(pq).(Content)
        PreloadContent(content)
    }
}

func main() {  
    var wg sync.WaitGroup
    pq := &PriorityQueue{}
    heap.Init(pq)
    
    // Start preload workers
    for i := 0; i < maxPreloadWorkers; i++ {  
        wg.Add(1)
        go preloadWorker(pq, &wg)
    }

    // Simulate traffic analysis and add content to the priority queue
    for {  
        regions := []string{"North America", "Europe", "Asia"}
        for _, region := range regions {  
            priority := rand.Intn(5) + 1
            content := Content{ID: fmt.Sprintf("Content-%d", rand.Intn(100)), Region: region, Priority: priority}
            heap.Push(pq, content)
            fmt.Printf("Added content %s to preload queue with priority %d\n", content.ID, content.Priority)
        }
        time.Sleep(5 * time.Second)
    }
    // Wait for all preload workers to complete (never happens in this infinite loop example)
    wg.Wait()
}