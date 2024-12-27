package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func doSomethingWithResource() {
	// Acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// Perform operations on the resource
	fmt.Println("Doing something with the resource.")
}

func main() {
	doSomethingWithResource()
	// The mutex will be unlocked automatically when doSomethingWithResource() returns
}