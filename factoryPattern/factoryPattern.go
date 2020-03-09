package main

import (
	"fmt"
	"github.com/krnblni/GoingGo/factoryPattern/vehicles"
)

func main() {
	// 0 - BIKE
	// 1 - CAR
	// 2 - TRUCK

	myVehicle, err := vehicles.CreateVehicle(2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Vehicle created")
	myVehicle.Start()
	fmt.Println(myVehicle.GetPurpose())
}