package vehicles

import (
	"fmt"
)

type Car struct {
	noOfWheels int
	purpose string
	isStarted bool
}

func NewCar() *Car {
	return &Car{
		noOfWheels: 4,
		purpose: "I have 4 wheels, that's it!",
		isStarted: false,
	}
}

func (car *Car) Start() {
	fmt.Println("Car started! :p")
	car.isStarted = true
}

func (car *Car) GetPurpose() string {
	return car.purpose
}