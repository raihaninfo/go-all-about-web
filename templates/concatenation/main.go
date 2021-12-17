package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Templating with concatenation
func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>` + name + `</h1>
	
	</body>
	</html>
	
	`)

	file, err := os.Create("index.gohtml")
	if err != nil {
		log.Fatal("error", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(str))
}
