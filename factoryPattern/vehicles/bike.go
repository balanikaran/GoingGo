package vehicles

import (
	"fmt"
)

type Bike struct {
	noOfWheels int
	purpose string
	isStarted bool
}

func NewBike() *Bike {
	return &Bike{
		noOfWheels: 2,
		purpose: "I have 2 wheels, that's it!",
		isStarted: false,
	}
}

func (bike *Bike) Start() {
	fmt.Println("Bike started! :p")
	bike.isStarted = true
}

func (bike *Bike) GetPurpose() string {
	return bike.purpose
}