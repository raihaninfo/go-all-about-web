package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	cuntry:= map[string]string{
		"Dhaka":"bangladesh",
		"Delhi": "india",
		"Kabul": "Afganistan",
		"Kathmandu": "Nepal",
	}

	tem, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	err = tem.Execute(os.Stdout, cuntry)
	if err != nil {
		log.Fatal(err)
	}

}
