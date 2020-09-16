package basic

var har int = 5
var tax int = 2
var Basic int

func Calculation() (allowance int, deduction int) {
	allowance = (Basic * har) / 100
	deduction = (Basic * tax) / 100
	return
}
