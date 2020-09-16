package main

import "fmt"

func main() {
	var country map[int]string

	country = make(map[int]string)

	country[0] = "India"

	for i, j := range country {
		fmt.Println("Key: %d value: %s\n", i, j)
	}
}
