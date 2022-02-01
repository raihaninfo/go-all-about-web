package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8082", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Println(fname, lname)

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
