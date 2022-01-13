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
	bangladesh:= cuntry{
		Name: "Banglades",
		Capital: "Dhaka",
	}
	india:= cuntry{
		Name: "India",
		Capital: "Delhi",
	}
	nepal:= cuntry{
		Name: "Nepal",
		Capital: "Catmundu",
	}

	cuntry:= []cuntry{bangladesh, india, nepal}

	tem, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	err = tem.Execute(os.Stdout, cuntry)
	if err != nil {
		log.Fatal(err)
	}

}
