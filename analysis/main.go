package main

import (
	human "analysis/human"
	vehicle "analysis/vehicle"
	"fmt"
)

func main() {
	var bmw vehicle.Vehicle
	bmw = vehicle.Car("world top brand")

	var labour human.Human
	labour = human.Man("software developer")

	for i, j := range bmw.Structure() {
		fmt.Println(" " + j + "" + labour.Structure()[i])
	}
}
