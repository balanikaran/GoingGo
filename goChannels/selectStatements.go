package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func getSomeMessage(c chan string) {
	sleepTime := rand.Intn(50)
	fmt.Println("Message will be sent after ", strconv.Itoa(sleepTime), "seconds")
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- strconv.Itoa(sleepTime)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c1 := make(chan string)
	c2 := make(chan string)

	go getSomeMessage(c1)
	go getSomeMessage(c2)

	fmt.Println("Awaiting message!")
	select {
	case msg := <- c1:
		fmt.Println("Message from Channel 1: ", msg)
	case msg := <- c2:
		fmt.Println("Message from Channel 2: ", msg)
	case <- time.After(20 * time.Second):
		fmt.Println("Timed out!!!")
	}
}