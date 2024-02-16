package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var t = template.Must(template.ParseGlob("templates/*.html"))
	var x = map[string]interface{} {
		"AboutMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-700 cursor-pointer",
		"PortfolioButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
		"ContactMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
		"AboutMe": true,
		"Portfolio": false,
		"ContactMe": false,
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("template executed")
		t.ExecuteTemplate(w, "index.html", x)
	})

	http.HandleFunc("/clicked", func (w http.ResponseWriter, r *http.Request) {
		var buttonName = r.URL.Query().Get("button")
		var d map[string]interface{}

		if buttonName == "about-me"{
			d = map[string]interface{} {
				"AboutMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-700 cursor-pointer",
				"PortfolioButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"ContactMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"AboutMe": true,
				"Portfolio": false,
				"ContactMe": false,
			}
		} else if buttonName == "portfolio"{
			d = map[string]interface{} {
				"AboutMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"PortfolioButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-700 cursor-pointer",
				"ContactMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"AboutMe": false,
				"Portfolio": true,
				"ContactMe": false,
			}
		} else if buttonName == "contact-me"{
			d = map[string]interface{} {
				"AboutMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"PortfolioButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-900 cursor-pointer",
				"ContactMeButtonClass":"p-3 hover:bg-cyan-700 hover:text-2xl rounded-t text-xl bg-cyan-700 cursor-pointer",
				"AboutMe": false,
				"Portfolio": false,
				"ContactMe": true,
			}
		}

		t.ExecuteTemplate(w, "index.html", d)
	})

	http.ListenAndServe(":8080", nil)
}