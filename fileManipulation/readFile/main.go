package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(text))
}
