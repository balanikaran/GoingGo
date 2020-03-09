package vehicles

import (
	"errors"
)

type Vehicles interface {
	Start()
	GetPurpose() string
}

// Vehicle types
const (
	BIKE = iota
	CAR
	TRUCK
)

func CreateVehicle(choice int) (Vehicles, error) {
	switch choice {
	case BIKE:
		return NewBike(), nil
	case CAR:
		return NewCar(), nil
	case TRUCK:
		return NewTruck(), nil
	default:
		return nil, errors.New("invalid vehicle type!")
	}
}