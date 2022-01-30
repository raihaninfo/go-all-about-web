package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Country string
	Age     string
}

func main() {
	birdJson := `{"Name": "Abdullah","Country":"Bangladesh","Age": "21 year"}`
	var p Person
	json.Unmarshal([]byte(birdJson), &p)
	fmt.Printf("Name: %s, Country: %s, Age: %s\n", p.Name, p.Country, p.Age)
}
