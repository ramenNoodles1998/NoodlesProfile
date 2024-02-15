package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Template struct {
	template *template.Template
	data map[string]interface{}
}

func main() {
	var templates []Template 
	renderer(&templates)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(templates); i++ { 
			var tmpl = templates[i]
			tmpl.template.Execute(w, tmpl.data)
		} 
	})

	http.ListenAndServe(":80", nil)
}

func renderer(templates *[]Template) {

	var tmpl, err = template.New("index").ParseFiles("index.html")

	if err != nil {
		fmt.Printf("Error parsing html. %#v", err.Error())
	} else {
		fmt.Printf("we gucci")
	}
	
	var template = Template {
		template: tmpl,
		data: map[string]interface{}{
			"Word": "Index",
		},
	}

	*templates = append(*templates, template)

	tmpl, err = tmpl.New("body").ParseFiles("body.html")

	if err != nil {
		fmt.Printf("Error parsing html. %#v", err.Error())
	} else {
		fmt.Printf("we gucci")
	}

	template = Template {
		template: tmpl,
		data: map[string]interface{}{
			"Word": "Body",
		},
	}

	*templates = append(*templates, template)
}
