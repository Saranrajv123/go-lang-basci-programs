package main

import "fmt"

func main() {
	var rows, temp int = 0, 1

	fmt.Print("Enter the number of rows ")
	fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" * ")
			temp++
		}
		fmt.Println("")
	}
}
