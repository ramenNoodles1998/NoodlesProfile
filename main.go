package main

import (
	"html/template"
	"net/http"
)

// type Template struct {
// 	template *template.Template
// 	data map[string]string
// }

// var TEMPLATE_MAP = map[string]map[string]string {
// 	"index": {
// 		"Word": "Index",
// 	},
// 	"body": {
// 		"Word": "Body",
// 	},
// }

func main() {
	var t = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index.gohtml", "Hello")
	})

	http.ListenAndServe(":8080", nil)
}