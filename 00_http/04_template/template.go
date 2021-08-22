package main

import (
	"net/http"
	"text/template"
)

// templating in Go
// golang automate escaping data for XSS prevention

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	// template can absorb all variant of data structure
	// from simple string or number, to array or struct

	// extract template from file
	// use must to auto panic if err
	tmpl := template.Must(template.ParseFiles("template.html"))
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// this data will fill template in template.html
		data := TodoPageData{
			PageTitle: "Todo Page",
			Todos: []Todo{
				{Title: "Todo 1", Done: true},
				{Title: "Todo 2", Done: false},
				{Title: "Todo 3", Done: false},
			},
		}
		// execute template use data and map to response writer
		tmpl.Execute(rw, data)
	})

	http.ListenAndServe(":8080", nil)
}
