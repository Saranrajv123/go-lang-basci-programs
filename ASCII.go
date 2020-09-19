package main

import "fmt"

func main() {
	// chVar := "T"
	// asciiCode := int(chVar)

	// fmt.Println("%s Ascii: %d ", string(chVar), asciiCode)

	// multiple two number

	var numb = []int{51, 4, 5, 6, 24}
	total := 0

	for _, num := range numb {
		total += num
	}

	avg := total / len(numb)
	fmt.Println("Avg :", avg)
}
