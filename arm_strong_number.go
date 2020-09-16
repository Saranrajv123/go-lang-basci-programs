package main

import "fmt"

func main() {
	var number, remainder, temp int
	result := 0

	fmt.Print("Enter the number ")
	fmt.Scanln(&number)

	temp = number

	for {
		remainder = temp % 10
		result += remainder * remainder * remainder
		temp /= 10

		if temp == 0 {
			break
		}
	}

	if number == result {
		fmt.Printf("%d arm strong number ", number)
	} else {
		fmt.Printf("%d Not arm strong number", number)

	}
}
