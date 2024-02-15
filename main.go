package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var t = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("template executed")
		t.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/clicked", func (w http.ResponseWriter, r *http.Request) {
		//TODO: add query params for each button.
		var htmlStr = "<div>Hello</div>"
		var tmpl, _ = template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}