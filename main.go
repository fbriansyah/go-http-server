package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/cat/", http.HandlerFunc(cat))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func errorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	errorHandler(res, err)
}

func cat(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "cat.gohtml", nil)
	errorHandler(res, err)
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "me.gohtml", "Febriansyah")
	errorHandler(res, err)
}
