package main

import "fmt"

func main() {
	var num [100]int
	var temp, sum, avg int

	fmt.Print("ENter the number of element: ")
	fmt.Scanln(&temp)

	for i := 0; i < temp; i++ {
		fmt.Print("ENter the number: ")
		fmt.Scanln(&num[i])
		sum += num[i]
	}

	avg = sum / temp
	fmt.Printf("The average of %d numbers: %d ", temp, avg)
}
