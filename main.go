package main

import (
	"fmt"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {

	var rect1 = rectangle{10, 20, "black"}
	var rect2 = rectangle{length: 10, breadth: 20, color: "black"}

	if rect1 == rect2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	rect3 := &rectangle{}
	rect4 := new(rectangle)

	if rect3 == rect4 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
