package father

import (
	"fmt"
)

func init() {
	fmt.Println("father pacakge init")
}

type Father struct {
	Name string
}

var Value = 10

func PrintData() {
	fmt.Println(Value)
}

func (f Father) Data(name string) string {
	f.Name = "Father :" + name
	return f.Name
}
