package main

import (
	"fmt"
	"net/http"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code want in this func")
}

func main() {
	var d anything
	http.ListenAndServe(":8082", d)
}
