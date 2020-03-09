package main

import (
	"fmt"
	cardoors "github.com/krnblni/GoingGo/builderPattern/carDoors"
)

func main() {
	carDoors := cardoors.NewCarDoorsBuilder().AddLeft().Build()
	carDoors2 := cardoors.NewCarDoorsBuilder().AddFront().AddLeft().Build()
	fmt.Println("Value: ", *carDoors)
	fmt.Println("Value2: ", *carDoors2)
}
