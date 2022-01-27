package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}
func main() {
	p := person{
		Name: "Jone Due",
		Age:  56,
	}

	tem, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	err = tem.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}

}
