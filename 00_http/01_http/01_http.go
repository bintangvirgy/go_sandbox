package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// process request
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome to web")
	})

	// serving static asset
	// # point to directory
	fs := http.FileServer(http.Dir("static/"))
	// # handle access from http to directory
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// listen connection
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
