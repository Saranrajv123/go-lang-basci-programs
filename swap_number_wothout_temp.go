package main

import "fmt"

func main() {
	var first, second, sum int

	fmt.Print("Enter 1st number: ")
	fmt.Scanln(&first)

	fmt.Print("enter 2nd number: ")
	fmt.Scanln(&second)

	// first = second
	// second = first
	first, second = second, first
	// first = first

	fmt.Println("first ", first)
	fmt.Println("second ", second)

	// sum program

	for i := 0; i <= first; i++ {
		sum += i
	}

	fmt.Println("sum ", sum)

}
