package main

import (
	"fmt"
	"net/http"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Raihan", "This is From Raihan")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "<h1>This is header</h1>")
}

func main() {
	var d anything
	http.ListenAndServe(":8082", d)
}
