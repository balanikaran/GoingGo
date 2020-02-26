package main

import (
	"strconv"
	"fmt"
)

func sendMultipleMessages(c chan string) {
	for i := 1; i <= 5; i++ {
		c <- strconv.Itoa(i)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go sendMultipleMessages(c)

	for s := range c {
		fmt.Println("Message: ", s)
	}

	_, isChannelOpen := <- c
	fmt.Println("isChannelOpen: ", isChannelOpen)
}