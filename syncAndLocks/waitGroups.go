package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			// do some work
			time.Sleep(2 * time.Second)
			fmt.Println("Work done: ", i)
		}(i)
	}

	wg.Wait()
}