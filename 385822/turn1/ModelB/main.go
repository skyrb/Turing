package main  
import (  
	"context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)
// Task represents a task to be scheduled
type Task struct {
	id    int
	fn    func()
	priority int // Lower value means higher priority
}
// Scheduler is a custom scheduler for managing Goroutines
type Scheduler struct {
	queue chan Task
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	m      sync.Mutex
	active atomic.Int32
}
// NewScheduler creates a new Scheduler instance
func NewScheduler() *Scheduler {
	sctx, cancel := context.WithCancel(context.Background())
	s := &Scheduler{
		queue: make(chan Task),
		ctx:   sctx,
		cancel: cancel,
	}
	s.wg.Add(1)
	go s.run()
	return s
}
// Run is a loop that runs indefinitely and picks up the tasks from the queue
func (s *Scheduler) run() {
	defer s.wg.Done()
	runtime.Gosched() // Give the main G the opportunity to schedule work
	for {
		select {
		case <-s.ctx.Done():
			return
		case task := <-s.queue:
			s.m.Lock()
			defer s.m.Unlock()
			s.scheduleTask(task)
		}
	}
}
// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	s.cancel()
	s.wg.Wait()
}
// AddTask enqueues a task with the specified priority
func (s *Scheduler) AddTask(task Task) {
	s.queue <- task
}
func (s *Scheduler) scheduleTask(task Task) {
	switch task.priority {
	case 0:
		go s.highPriority(task)
	default:
		go s.lowPriority(task)
	}
}
func (s *Scheduler) highPriority(task Task) {
	log.Printf("Running high priority task: Task ID: %d", task.id)
	s.runTask(task)
}
func (s *Scheduler) lowPriority(task Task) {
	log.Printf("Running low priority task: Task ID: %d", task.id)
	s.runTask(task)
}
func (s *Scheduler) runTask(task Task) {
	start := time.Now()
	task.fn()
	end := time.Now()
	duration := end.Sub(start)
	log.Printf("Task ID: %d finished in %s", task.id, duration)
}
func main() {
	scheduler := NewScheduler()

	scheduler.AddTask(Task{
		id:    1,
		fn:    func() { time.Sleep(2 * time.Second) }, // Low priority, takes 2s
		priority: 1,
	})

	scheduler.AddTask(Task{
		id:    2,
		fn:    func() { time.Sleep(500 * time.Millisecond) }, // High priority, takes 500ms
		priority: 0,
	})

	scheduler.AddTask(Task{
		id:    3,
		fn:    func() { time.Sleep(1 * time.Second) }, // Low priority, takes 1s
		priority: 1,
	})

	<-time.After(3 * time.Second)
	scheduler.Stop()
}