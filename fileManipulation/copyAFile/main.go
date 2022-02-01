package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	mainFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	newFile, err := os.Create("compFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	copy, err := io.Copy(mainFile, newFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(copy)
}
