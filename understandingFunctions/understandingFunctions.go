package main

import (
	"fmt"
)

// functionA {no arguments, no return type}
func functionA(){
	fmt.Println("message from function A")
}

// functionB {argument, no return type}
func functionB(message string){
	fmt.Println("message from function B: ", message)
}

// functionC {no argument, 1 return type}
func functionC() (int) {
	return 2
}

// functionD {2 arguments, 3 return types with names}
func functionD(a, b int) (addition int, subtraction int, multiplication int) {
	return a+b, a-b, a*b
}

// functionE variation of functionD
func functionE(a, b int) (addition int, subtraction int, multiplication int) {
	addition = a + b
	subtraction = a - b
	multiplication = a * b
	return
}

// functionF adds a 3 times {uses function passed as an argument}
func functionF(a int, fn func(b, c int) int) int {
	return a + fn(a, a)
}

// increment function uses FUNCTION CLOSURE
// basically, a function returning another function
func increment() func() int {
	// initializing i
	i := 0
	return func() int {
		// increment i
		i ++
		return i
	}
}

// functionG is basic function which
// takes reference argument
func functionG(x *int) {
	*x = *x * 2
}

func main(){
	functionA()
	functionB("LOL from main()")
	fmt.Println("functionC: ", functionC())
	fmt.Println(functionD(2, 3))
	fmt.Println(functionE(2, 3))

	add, sub, mul := functionE(3, 4)
	fmt.Println(add, sub, mul)

	// taking function as a variable
	addFn := func (a, b int) (int) {
		return a + b
	}
	fmt.Println(functionF(5, addFn))

	// calling this line initializes the value of i
	// and a function is returned
	inc := increment()
	// calling inc() repetedely calls the insider function
	fmt.Println(inc(), inc(), inc())

	// same applies here too...
	inc2 := increment()
	fmt.Println(inc2(), inc2())

	x := 10
	fmt.Println("before function call x = ", x)
	// this is nothing but passing the memory location of variable x
	// so that it's value can be altered or used by any other function
	// and reflecting changes back here in this function
	functionG(&x)
	fmt.Println("after function call x = ", x)
}