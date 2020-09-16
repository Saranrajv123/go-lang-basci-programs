package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU, "numcpu")

	intSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("last element : ", len(intSlice)-1)
	fmt.Println("last element : ", intSlice[2:4])
	fmt.Println("last element : ", intSlice[:len(intSlice)-1])
	fmt.Println("last element : ", intSlice[1:])

	newValue := [3]string{"new", "go", "lang"}
	newv1 := newValue
	newv2 := &newValue
	fmt.Println("newValue: ", newv1)
	fmt.Println("*newValue: ", newv2)

	newValue[0] = "learning"
	fmt.Println("newValue", newValue)
	fmt.Println("newValue: ", newv1)
	fmt.Println("*newValue: ", *newv2)

}
