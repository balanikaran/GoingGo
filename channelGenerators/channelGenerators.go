package main

import (
	"fmt"
	"time"
)

func tickerCounter(ticker *time.Ticker) chan bool {

	done := make(chan bool)
	go func() {
		i := 0
		Loop:
		for {
			select {
			case t := <- ticker.C:
				i++
				fmt.Println("Count", i, " at: ", t)
			case <- done:
				fmt.Println("Done signal")
				break Loop
			}
		}
		fmt.Println("Exiting ticker counter...")
	}()
	return done
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := tickerCounter(ticker)
	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting main...")
}
