package main

import "fmt"

func main() {
	var first, second, third, sum int

	fmt.Print("Enter the first Number: ")
	fmt.Scanln(&first)

	fmt.Print("Enter the first Number: ")
	fmt.Scanln(&second)

	fmt.Print("Enter the first Number: ")
	fmt.Scanln(&third)

	if first == second && second == third {
		fmt.Println("equal value")
	} else {
		sum = first + second + third
	}

	fmt.Println(sum, "sum")
}
