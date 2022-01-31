package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println(file)
}
