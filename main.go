package main

import (
	"fmt"
	"html/template"
	"net/http"
)


type Data struct {
	NavButtonClass func(string) string 
	TemplateName string
}

var data Data 

func main() {
	var t = template.Must(template.ParseGlob("templates/*.html"))
	data = Data {
		NavButtonClass: NavButtonClass,
		TemplateName: "about-me",
	}	
	// data = map[string]interface{} {
	// 	"NavButtonClass": NavButtonClass,
	// 	"GetTemplate": GetTemplate,
	// 	"TemplateName": "about-me",
	// }
	//TODO: First fix button classes to change.
	//don't re render whole page just changing template part.


	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("template executed")
		t.ExecuteTemplate(w, "index.html", data)
	})

	http.HandleFunc("/clicked", func (w http.ResponseWriter, r *http.Request) {
		data.TemplateName = r.URL.Query().Get("button")

		t.ExecuteTemplate(w, "index.html", data)
	})

	http.ListenAndServe(":8080", nil)
}

func NavButtonClass(buttonName string) string {
	if(buttonName == data.TemplateName) {
		return "p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-700 cursor-pointer"
	}
	return "p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer"
}
