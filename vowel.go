package main

import "fmt"

func isVowel(ch rune) {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func isVowelAnotherWay(ch rune) {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("isVowelAnotherWay", true)
	default:
		fmt.Println("isVowelAnotherWay", false)
	}

}

func calcuateVowelCount(vowel string) {

	// str := "go programming"
	count := 0

	for _, value := range vowel {
		switch value {
		case 'a', 'e', 'i', 'o', 'u':
			count += 1
		}
	}

	fmt.Println("count ", count)

}

func main() {
	// var ch rune

	// fmt.Println("Enter the char: ")
	// fmt.Scanln(&ch)

	isVowel('a')
	isVowelAnotherWay('f')
	calcuateVowelCount("go programming")

}
