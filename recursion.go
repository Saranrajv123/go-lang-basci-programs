package main

import "fmt"

var count int = 0

func recursionFunc() {
	fmt.Println("Runing...")
	// recursionFunc()
}

func recursionCountMethod() {
	count++
	if count <= 10 {
		fmt.Println(count)
		recursionCountMethod()
	}
}

func main() {
	recursionFunc()
	recursionCountMethod()
}
