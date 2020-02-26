package main

import (
	"fmt"
	"time"
)

func doPanic() error {
	// simulate panic creation :p
	time.Sleep(2 * time.Second)

	panic("Something is wrogn!")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic was recovered!", err)
			fmt.Println("Yeah it was the spelling! :p")
		}
	}()

	if err:= doPanic(); err != nil {
		// some error occured
		fmt.Println("Error occured!")
	}

}