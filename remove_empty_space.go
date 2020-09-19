package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  This is for Example "
	fmt.Println("text :", strings.Replace(str, "", "", -1))
	fmt.Println("text :", strings.TrimSpace(str))
}
