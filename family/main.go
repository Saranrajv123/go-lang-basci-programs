package main

import (
	parent "family/father"
	"fmt"
)

func main() {
	father := new(parent.Father)
	fmt.Println(father.Data("Father "))

	fmt.Printf("%q\n", father.Data("father"))

	// c := &parent.Value
	// fmt.Println(c)
	// fmt.Println(*c)

	// *c++

	a := 3
	b := &a
	*b++
	fmt.Println(*b)
	parent.PrintData()

}
