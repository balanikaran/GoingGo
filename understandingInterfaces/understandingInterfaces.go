package main

import (
	"fmt"
)

// MyTypes - may have these methods
type MyTypes interface {
	SetValue(value int)
	GetValue() int
}


// MySimpleType - custom type
// This is child of MyType because it implements both of its methods
type MySimpleType struct {
	value int
}

// SetValue of variable
func (variable *MySimpleType) SetValue(value int) {
	variable.value = value
}

// GetValue of variable
func (variable *MySimpleType) GetValue() int {
	return variable.value
}

// NewMySimpleType Create new MySimpleType variable
func NewMySimpleType() *MySimpleType {
	return new(MySimpleType)
}


// MySquareType - custom type
// This is child of MyType because it implements both of its methods
type MySquareType struct {
	squareValue int
}

// SetValue of variable
func (variable *MySquareType) SetValue(value int){
	variable.squareValue = value * value
}

// GetValue of variable
func (variable *MySquareType) GetValue() int {
	return variable.squareValue
}

// NewMySquareType Create new MySquareType variable
func NewMySquareType() *MySquareType {
	return new(MySquareType)
}

func main() {

	var myType MyTypes

	myType = NewMySimpleType()
	myType.SetValue(5)
	fmt.Println("Value of myType [MySimpleType]: ", myType.GetValue())

	myType = NewMySquareType()
	myType.SetValue(5)
	fmt.Println("Value of myType [MySquareType]: ", myType.GetValue())

}