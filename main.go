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
	var t = template.Must(template.ParseGlob("templates/*.gohtml"))
	data = Data {
		NavButtonClass: NavButtonClass,
		TemplateName: "about-me",
	}	

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("template executed")
		t.ExecuteTemplate(w, "index.gohtml", data)
	})

	http.HandleFunc("/clicked", func (w http.ResponseWriter, r *http.Request) {
		data.TemplateName = r.URL.Query().Get("button")

		t.ExecuteTemplate(w, "index.gohtml", data)
	})

	http.ListenAndServe(":3000", nil)
}

func NavButtonClass(buttonName string) string {
	if(buttonName == data.TemplateName) {
		return "p-3 rounded-t text-2xl bg-cyan-700 cursor-pointer"
	}
	return "p-3 hover:bg-cyan-700 hover:text-2xl  hover:text-slate-100 text-xl text-slate-500 cursor-pointer"
}
