package main  
import (  
    "container/heap"
    "fmt"
    "sync"
    "time"
)

type Notification struct {  
    Type     string
    Data     interface{}
    Priority int // Priority level for the notification
}

// PriorityQueue implements heap.Interface and holds Notification items.
type PriorityQueue []*Notification

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want notifications with higher priority to have lower index, so we use < here.
    return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Notification)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

var (  
    subscriptions       map[int][]chan Notification
    groupSubscriptions map[int]map[int]bool  // Map of group ID to a set of user IDs
    subscriptionsMutex  sync.Mutex
    priorityQueue      PriorityQueue
    priorityQueueMutex sync.Mutex
)

func init() {  
    subscriptions = make(map[int][]chan Notification)
    groupSubscriptions = make(map[int]map[int]bool)
    heap.Init(&priorityQueue)
}

// ... (Same functions as before for user subscriptions and unsubscribe)

// Subscribe to a group
func subscribeGroup(userID int, groupID int) {
    subscriptionsMutex.Lock()
    defer subscriptionsMutex.Unlock()

    _, ok := groupSubscriptions[groupID]
    if !ok {
        groupSubscriptions[groupID] = make(map[int]bool)
    }
    groupSubscriptions[groupID][userID] = true
}

// Unsubscribe from a group
func unsubscribeGroup(userID int, groupID int) {
    subscriptionsMutex.Lock()
    defer subscriptionsMutex.Unlock()

    delete(groupSubscriptions[groupID], userID)
}

// Notify a single user with a notification
func notifyUser(userID int, notification Notification) {
    subscriptionsMutex.Lock()
    defer subscriptionsMutex.Unlock()

    for _, ch := range subscriptions[userID] {
        go func(ch chan Notification) {
            ch <- notification
        }(ch)
    }
}

// Notify all users in a group with a notification
func notifyGroup(groupID int, notification Notification) {
    subscriptionsMutex.Lock()
    defer subscriptionsMutex.Unlock()

    for userID := range groupSubscriptions[groupID] {
        notifyUser(userID, notification)
    }
}

// Send a targeted broadcast to a specific set of users (by ID)
func targetedBroadcast(userIDs []int, notification Notification) {
    for _, userID := range userIDs {
        notifyUser(userID, notification)
    }
}

// Function to add a notification to the priority queue
func addNotificationToQueue(notification *Notification) {
    priorityQueueMutex.Lock()
    defer priorityQueueMutex.Unlock()

    heap.Push(&priorityQueue, notification)
}

// Function to process notifications from the priority queue
func processNotifications() {