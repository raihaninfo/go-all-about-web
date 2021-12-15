package main

import "fmt"

type persion struct {
	name string
	age  int
}

func (p persion) info() {
	fmt.Println(p.name, "is", p.age, "year old")
}

func main() {
	p1 := persion{
		"raihan",
		20,
	}
	fmt.Println(p1)
	p1.info()
}
