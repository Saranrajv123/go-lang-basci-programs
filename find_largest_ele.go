package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp int
	var num [100]float64

	fmt.Println(strings.Compare("India", "Indian"))

	fmt.Print("Element the number: ")
	fmt.Scanln(&temp)

	fmt.Println("num ", num)

	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&num[i])
	}

	for j := 0; j < temp; j++ {
		if num[0] < num[j] {
			num[0] = num[j]
		}
	}

	fmt.Println("num[0] ", num[0])
}
