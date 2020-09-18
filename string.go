package main

import "fmt"

func main() {
	var str1 string

	fmt.Print("Enter the String :")
	fmt.Scanln(&str1)

	if str1[0:2] == "Is" {
		fmt.Println("equal")
	} else {
		fmt.Println("nope")
	}
}
