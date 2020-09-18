package main

import "fmt"

func listCount(numbers []int) {
	count := 0

	for _, num := range numbers {
		if num == 4 {
			count += 1
		}
	}

	fmt.Println("Count: ", count)
}

func printUniqueValue(listCount []int) {
	dict := make(map[int]int)
	fmt.Println("dict", dict)

	for _, value := range listCount {
		fmt.Println(dict[value], "value")
		dict[value] += 1
	}

	fmt.Println(dict)
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 4}
	listCount(num)
	printUniqueValue(num)
}
