// Understanding package text/template: parsing & executing templates
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tem, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	err = tem.Execute(os.Stdout, 34)
	if err != nil {
		log.Fatal(err)
	}

}
