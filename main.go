package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	var tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("Error parsing html. %#v", err.Error())
	} else {
		fmt.Printf("we gucci")
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)
}
