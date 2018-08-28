package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tmpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func httpErrorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		// open file

		f, h, err := req.FormFile("q")
		if err != nil {
			httpErrorHandler(w, err)
			return
		}
		f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			httpErrorHandler(w, err)
			return
		}
		s = string(bs)

		//store
		dst, err := os.Create(filepath.Join("./storage-fs/", h.Filename))
		if err != nil {
			httpErrorHandler(w, err)
			return
		}

		_, err = dst.Write(bs)
		if err != nil {
			httpErrorHandler(w, err)
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
