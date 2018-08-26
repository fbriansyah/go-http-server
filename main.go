package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type anything int

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>this is anything<h1>")
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	fmt.Println("Serve...")
	var an anything

	http.ListenAndServe(":8080", an)
}
