package main

import "fmt"

func factorialFunc(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorialFunc(num-1)
}

func main() {
	result := factorialFunc(5)
	fmt.Println("output :", result)
}
