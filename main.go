package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	template *template.Template
	data map[string]string
}

var TEMPLATE_MAP = map[string]map[string]string {
	"index": {
		"Word": "Index",
	},
	"body": {
		"Word": "Body",
	},
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

	http.ListenAndServe(":8080", nil)
}

func renderer(templates *[]Template) {
	for k, v := range TEMPLATE_MAP { 
		fmt.Println(k)
		fmt.Println(v)
		var tmpl, err = template.New(k).ParseFiles(k + ".html")

		if err != nil {
			fmt.Printf("Error parsing html. %#v\n", err.Error())
		} else {
			fmt.Println("we gucci")
		}
		
		var t = Template {
			template: tmpl,
			data: v,
		}

		*templates = append(*templates, t)
	}

	// var tmpl, err = template.New("index").ParseFiles("index.html")

	// if err != nil {
	// 	fmt.Printf("Error parsing html. %#v", err.Error())
	// } else {
	// 	fmt.Printf("we gucci")
	// }
	
	// var template = Template {
	// 	template: tmpl,
	// 	data: map[string]string{
	// 		"Word": "Index",
	// 	},
	// }

	// *templates = append(*templates, template)

	// tmpl, err = tmpl.New("body").ParseFiles("body.html")

	// if err != nil {
	// 	fmt.Printf("Error parsing html. %#v", err.Error())
	// } else {
	// 	fmt.Printf("we gucci")
	// }

	// template = Template {
	// 	template: tmpl,
	// 	data: map[string]string{
	// 		"Word": "Body",
	// 	},
	// }

	// *templates = append(*templates, template)
}

