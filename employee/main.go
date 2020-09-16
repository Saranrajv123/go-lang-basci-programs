package main

import (
	basic "employee/basic"
	gross "employee/basic/gross"
	"fmt"
)

func main() {
	basic.Basic = 10000
	fmt.Println(gross.GrossSalary())
}
