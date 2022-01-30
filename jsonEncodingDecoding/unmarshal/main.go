package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	birdJson := `{"Name": "pigeon","Age": "21 year"}`
	var p Person
	json.Unmarshal([]byte(birdJson), &p)
	fmt.Printf("Name: %s, Age: %v\n", p.Name, p.Age)
}
