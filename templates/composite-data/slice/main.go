package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	fruits := []string{"mango", "banana", "pineApple", "Apple"}

	tem, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	err = tem.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatal(err)
	}

}
