package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	fmt.Printf("%T\n", str)
	fmt.Println(strconv.Atoi(str))
	fmt.Println(strconv.ParseInt(str, 10, 34))
}
