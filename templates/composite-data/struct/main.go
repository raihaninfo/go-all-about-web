package main

import (
	"html/template"
	"log"
	"os"
)

type cuntry struct{
	Name string
	Capital string
}

func main() {
	cuntry:= cuntry{
		Name: "Banglades",
		Capital: "Dhaka",
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
