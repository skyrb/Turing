package main  
import (  
    "fmt"
    "sync"
    "time"
)
type ReadWriteQueue struct {  
    queue []string
    readersWg sync.WaitGroup
    writersWg sync.WaitGroup
    readerLock sync.RWMutex
    writerLock sync.Mutex
}
func NewReadWriteQueue() *ReadWriteQueue {  
    return &ReadWriteQueue{}
}
func (q *ReadWriteQueue) Write(item string) {  
    q.writerLock.Lock()
    defer q.writerLock.Unlock()
    q.queue = append(q.queue, item)
    q.readersWg.Broadcast()
}
func (q *ReadWriteQueue) Read() (string, bool) {  
    q.readerLock.Lock()
    defer q.readerLock.Unlock()
    q.readersWg.Add(1)
    defer q.readersWg.Done()
    for len(q.queue) == 0 {
        q.readersWg.Wait()
    }
    item := q.queue[0]
    q.queue = q.queue[1:]
    return item, true
}
func (q *ReadWriteQueue) StartReadGoroutines(count int, interval time.Duration) {  
    for i := 0; i < count; i++ {  
        go func() {  
            for {  
                item, ok := q.Read()  
                if !ok {  
                    break
                }  
                fmt.Println("Read:", item)  
                time.Sleep(interval)
            }
        }()
    }
}
func (q *ReadWriteQueue) WaitForCompletion() {  
    q.writersWg.Wait()
}
func main() {  
    queue := NewReadWriteQueue()
    queue.StartReadGoroutines(5, time.Second/2)
    go func() {  
        defer queue.writersWg.Done()
        for i := 0; i < 10; i++ {  
            queue.Write(fmt.Sprintf("Item %d", i))
            time.Sleep(time.Second / 4)
        }
    }()
    go func() {  
        defer queue.writersWg.Done()
        for i := 10; i < 20; i++ {  
            queue.Write(fmt.Sprintf("Item %d", i))
            time.Sleep(time.Second / 4)
        }
    }()
    queue.writersWg.Add(2)
    queue.WaitForCompletion()
    fmt.Println("All write operations completed.")
}
