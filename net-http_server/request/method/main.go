package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d anything
	http.ListenAndServe(":8082", d)
}
