package main

import (
	"html/template"
	"log"
	"net/http"
)

type anything int

var tpl *template.Template

func (m anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d anything
	http.ListenAndServe(":8082", d)
}
