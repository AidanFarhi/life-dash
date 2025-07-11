package main

import (
	"fmt"
	"html/template"
	"lifedash/handler"
	"net/http"
	"os"
)

func main() {

	tmpl, err := template.ParseFiles(
		"templates/pages/index.html",
		"templates/pages/login.html",
		"templates/partials/hub.html",
	)
	if err != nil {
		fmt.Println("error parsing templates:", err.Error())
		os.Exit(1)
	}

	indexHandler := handler.IndexHandler(tmpl)

	m := http.NewServeMux()
	m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	m.Handle("/", indexHandler)

	s := http.Server{
		Addr: ":1337",
		Handler: m,
	}
	
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println("error starting server:", err.Error())
		os.Exit(1)
	}
}