package main

import (
	"sync"
	"fmt"
)

func main() {
	var once sync.Once

	myOnceFunc := func() {
		fmt.Println("My once function was called!")
	}

	for i := 0; i < 10; i++ {
		once.Do(myOnceFunc)
	}
	
}