package main

import "fmt"

func getPerms(str string) []string {

	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1]
	remainder := str[1:]

	parms := getPerms(remainder)

	fmt.Println(parms, "perms")

	allParms := make([]string, 0)

	for _, v := range parms {
		fmt.Println(v, "v=======")
		for i := 0; i <= len(v); i++ {
			newParam := insertAt(i, current, v)
			allParms = append(allParms, newParam)
		}
	}

	fmt.Println(allParms, "allparams")

	return allParms
}

func insertAt(i int, current string, v string) string {
	fmt.Println("parameter", i, current)
	fmt.Println("value", v)

	start := v[0:i]
	end := v[i:len(v)]

	fmt.Println("start====", start+current+end)
	return start + current + end

}

func main() {
	fmt.Println("permutation: ", getPerms("abc"))
}
