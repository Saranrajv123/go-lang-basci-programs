package main

import "fmt"

func leap_year(year int) bool {
	if year%2 == 0 {
		if year%400 == 0 {
			if year%100 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}

}

func main() {
	year := 2000
	fmt.Println(leap_year(year))
}
