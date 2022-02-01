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

	write, err := file.Write([]byte("Test file"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(write)
}
