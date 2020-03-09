package vehicles

import (
	"fmt"
)

type Truck struct {
	noOfWheels int
	purpose string
	isStarted bool
}

func NewTruck() *Truck {
	return &Truck{
		noOfWheels: 4,
		purpose: "I have 4 wheels and I can carry huge weight!",
		isStarted: false,
	}
}

func (truck *Truck) Start() {
	fmt.Println("Truck started! :p")
	truck.isStarted = true
}

func (truck *Truck) GetPurpose() string {
	return truck.purpose
}