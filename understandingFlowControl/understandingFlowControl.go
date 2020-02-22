package main

import (
	"fmt"
)

// checkPolarity1 uses if-else block to check polarity
func checkPolarity1(val int) string {
	if val > 0 {
		return "positive"
	} else if val < 0 {
		return "negative"
	} else {
		return "zero"
	}
}

// checkPolarity2 uses switch block to check polarity
func checkPolarity2(val int) string {
	switch {
	case val > 0:
		return "positive"
	case val < 0:
		return "negative"
	default:
		return "zero"
	}
}

// simpleForLoop implements a simple for loop. wow! :p
func simpleForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

// whileLoop implements a while loop like structure using for keyword
func whileLoop(){
	// initialize
	i := 0
	for i < 10 { // condition
		fmt.Print(i, " ")
		i++ // updation
	}
}

// infiniteLoop using for keyword
func infiniteLoop(){
	for {
		// my infinite code :)
	}
}

func understandingDeferKeyword(){
	// defer is called when a function completes executing
	// called in a LAST IN FIRST OUT MANNER
	fmt.Println("Counting start...")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Exiting, now defer will be called")
}

func main() {
	fmt.Println("-2 is: ", checkPolarity1(-2))
	fmt.Println("5 is: ", checkPolarity2(5))
	simpleForLoop()
	whileLoop()
	// not calling this...
	// infiniteLoop()
	understandingDeferKeyword()
}