package main

import (
	"fmt"
	"time"
	"sync"
)

type safeMapCounter struct {
	m map[int]int 
	// embedding :)
	sync.RWMutex
}

func writers(sm *safeMapCounter, n int) {
	for i := 0; i < n; i ++{
		sm.Lock()
		sm.m[i] = i * 100
		sm.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func readers(sm *safeMapCounter, key int) {
	// this loop will end when main() ends
	for {
		sm.RLock()
		value := sm.m[key]
		sm.RUnlock()
		fmt.Println("Value: ", value)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	mySm := safeMapCounter{m: make(map[int]int)}
	go writers(&mySm, 20)
	go readers(&mySm, 4)
	go readers(&mySm, 7)
	time.Sleep(20 * time.Second)
}