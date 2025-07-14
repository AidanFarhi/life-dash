package main

import (
	"fmt"
	"html/template"
	"lifedash/handler"
	"lifedash/service"
	"net/http"
	"os"
	"path/filepath"
)

func parseTemplates() (*template.Template, error) {
	fileNames, err := filepath.Glob("templates/**/*.html")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.ParseFiles(fileNames...)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func main() {

	tmpl, err := parseTemplates()
	if err != nil {
		fmt.Println("error parsing templates:", err.Error())
		os.Exit(1)
	}

	authService := service.NewAuthService()

	indexHandler := handler.IndexHandler(authService, tmpl)
	expenseHandler := handler.ExpensesHandler(tmpl)
	hubHandler := handler.HubHandler(tmpl)

	m := http.NewServeMux()
	m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	m.Handle("/", indexHandler)
	m.Handle("GET /expenses", expenseHandler)
	m.Handle("GET /hub", hubHandler)

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