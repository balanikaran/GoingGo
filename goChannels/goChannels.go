package main

import (
	"time"
	"fmt"
)

func driver() {
	c := make(chan bool)
	go waitForDriver(c, "LOL")

	fmt.Println("hello from driver, sending msg in 5 seconds")

	// we can now send msg to waitForDriver()
	time.Sleep(5 * time.Second)
	c <- true

	// we wait until waitForDriver() completes and sends back a bool value
	<- c
}

func waitForDriver(c chan bool, msg string) {
	if b := <- c; b {
		fmt.Println("driver sent true, msg: ", msg)
	}
	c <- true
}

func main() {
	driver()
}