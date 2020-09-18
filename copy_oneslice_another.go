package main

import "fmt"

func main() {
	slc1 := []int{22, 34, 55, 67, 78, 2, 29, 20}
	var slc2 []int
	var slc3 []int

	cp1 := append(slc2, slc1...)
	cp2 := copy(slc1, slc3)

	fmt.Println("cp1", cp1)
	fmt.Println("cp2", cp2)

}
