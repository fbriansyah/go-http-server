package main

import (
	"fmt"
	"net/http"
)

type anything int

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>this is anything<h1>")
}

func main() {
	fmt.Println("Serve...")
	var an anything

	http.ListenAndServe(":8080", an)
}
