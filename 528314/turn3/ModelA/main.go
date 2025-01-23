package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Task represents a task with a name, priority, and deadline.
type Task struct {
	TaskName string
	Priority int
	Deadline time.Time
	index    int // The index of the task in the heap; maintained by the heap.Interface methods.
}

// PriorityQueue implements a priority queue for tasks.
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want a max-heap by Priority.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes an element into the heap
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.index = n
	*pq = append(*pq, task)
}

// Pop removes and returns the minimum element (according to Less) from the heap
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // Avoid memory leak
	task.index = -1
	*pq = old[0 : n-1]
	return task
}

// update modifies the priority and Value of a Task in the queue.
func (pq *PriorityQueue) update(task *Task, priority int) {
	task.Priority = priority
	heap.Fix(pq, task.index)
}

// TaskScheduler manages the prioritized tasks with concurrency control.
type TaskScheduler struct {
	tasks PriorityQueue
	mu    sync.Mutex // Protects the task priority queue.
}

// NewTaskScheduler creates a new Task Scheduler
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{tasks: make(PriorityQueue, 0)}
}

// AddTask adds a task to the scheduler.
func (s *TaskScheduler) AddTask(taskName string, priority int, deadline time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task := &Task{
		TaskName: taskName,
		Priority: priority,
		Deadline: deadline,
	}
	heap.Push(&s.tasks, task)
}

// UpdatePriorityByDeadline recalculates priorities based on deadlines.
func (s *TaskScheduler) UpdatePriorityByDeadline() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for _, task := range s.tasks {
		daysToDeadline := task.Deadline.Sub(now).Hours() / 24
		if daysToDeadline < 0 {
			s.tasks.update(task, 1) // Overdue tasks
		} else if daysToDeadline < 2 {
			s.tasks.update(task, 5) // Due within 2 days
		} else if daysToDeadline < 5 {
			s.tasks.update(task, 3) // Due within 5 days
		}
	}
}

// CompleteTask removes the highest priority task
func (s *TaskScheduler) CompleteTask() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.tasks) > 0 {
		task := heap.Pop(&s.tasks).(*Task)
		fmt.Printf("Completed Task: %s\n", task.TaskName)
	}
}

func main() {
	scheduler := NewTaskScheduler()
	scheduler.AddTask("Task 1", 5, time.Date(2023, 11, 5, 0, 0, 0, 0, time.UTC))
	scheduler.AddTask("Task 2", 3, time.Date(2023, 11, 10, 0, 0, 0, 0, time.UTC))
	scheduler.AddTask("Task 3", 4, time.Date(2023, 11, 3, 0, 0, 0, 0, time.UTC))

	// Simulate concurrent updates
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scheduler.UpdatePriorityByDeadline()
		}()
	}
	wg.Wait()

	// Complete a task
	scheduler.CompleteTask()

	fmt.Printf("Remaining Tasks: %v\n", scheduler.tasks)
}