package main

import (
	"sync"
	"time"
	"math/rand"
	"log"
)

type AnnotatedMap struct {
	sync.RWMutex
	m map[int]int
}

func (am *AnnotatedMap) Set(key int, value int) {
	am.Lock()
	defer am.Unlock()
	am.m[key] = value
}

func (am *AnnotatedMap) Get(key int) (int, bool) {
	am.RLock()
	defer am.RUnlock()
	return am.m[key], am.m[key] != 0
}

func (am *AnnotatedMap) Delete(key int) {
	am.Lock()
	defer am.Unlock()
	delete(am.m, key)
}


func main() {
	am := AnnotatedMap{}

	go concurrentMapOperations(&am)
	go concurrentMapOperations(&am)
	go concurrentMapOperations(&am)

	time.Sleep(1 * time.Second)
}

func concurrentMapOperations(am *AnnotatedMap) {
	for {
		key := rand.Intn(100)
		value := rand.Intn(100)

		am.Set(key, value)

		val, _ := am.Get(key)
		log.Printf("Key %d: Value %d\n", key, val)

		time.Sleep(time.Millisecond * 100)
	}
}