package gross

import (
	basic "employee/basic"
	"fmt"
)

func GrossSalary() int {
	a, t := basic.Calculation()
	fmt.Println("a", a, t)
	return ((basic.Basic + a) - t)

}
