package main

import (
	"fmt"
	"os"
)

func main() {
	//rename()
	delete()
}

func rename() {
	err := os.Rename("file.txt", "file2.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func delete() {
	err := os.Remove("newfile.txt")
	if err != nil {
		fmt.Println(err)
	}
}
