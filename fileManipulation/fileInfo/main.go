package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Size())
}
