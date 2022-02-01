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
	defer mainFile.Close()
	newFile, err := os.Create("compFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	copy, err := io.Copy(newFile, mainFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(copy)
}
