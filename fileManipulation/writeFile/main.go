package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("file.txt", 333, 666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileText := "This is dymmy text"

	write, err := file.Write([]byte(fileText))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(write)
}
