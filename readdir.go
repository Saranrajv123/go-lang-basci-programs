package main

import (
	"fmt"
	"os"
)

func readerDir() {
	file, err := os.Open(".")

	if err != nil {
		fmt.Println(err)
	}

	fileList, _ := file.Readdir(0)
	for _, file := range fileList {
		fmt.Println(file.Name(), file.Size())
	}
}
func main() {
	readerDir()
}
