package main

import (
	"time"
	"fmt"
	"sync"
)

type safeCounter struct {
	counter int
	// embedding is done here :)
	sync.Mutex
}

func (sf *safeCounter) increment() {
	sf.Lock()
	sf.counter ++
	sf.Unlock()
}

func (sf *safeCounter) decrement() {
	sf.Lock()
	sf.counter --
	sf.Unlock()
}

func (sf *safeCounter) getCounter() int {
	sf.Lock()
	value := sf.counter
	sf.Unlock()
	return value
}

func main() {
	mySafeCounter := new(safeCounter)

	for i := 0; i < 100; i ++ {
		go mySafeCounter.increment()
		go mySafeCounter.decrement()
	}

	fmt.Println("Counter value: ", mySafeCounter.getCounter())
}