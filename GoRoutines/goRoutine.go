package main

import (
	"fmt"
	"time"
)

func greetAfterDelay(msg string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("message from go routine: ", msg)
}

func routinesCanUseSameMemory(){
	value := 5

	go func () {
		time.Sleep(5 * time.Second)
		fmt.Println("Value: ", value)
	} ()
	
	fmt.Println("Value: ", value)
	value = 2 * value // updated
	time.Sleep(6 * time.Second)

}

func main() {
	delay := 5
	msg := "alpha"

	go greetAfterDelay(msg, delay)

	fmt.Println("this is main")
	time.Sleep(time.Duration(delay + 1) * time.Second)

	routinesCanUseSameMemory()
}