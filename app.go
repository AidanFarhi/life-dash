package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"lifedash/handler"
	"lifedash/repo"
	"lifedash/service"
	"net/http"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./db/lifedash.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func configureDb(db *sql.DB) error {
	_, err := db.Exec("PRAGMA foreign_keys=ON;")
	return err
}

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

func handleError(operation string, err error) {
	if err != nil {
		fmt.Println("error "+operation+":", err.Error())
		os.Exit(1)
	}
}

func main() {

	db, err := getDB()
	handleError("getting db", err)

	err = configureDb(db)
	handleError("configuring db", err)

	tmpl, err := parseTemplates()
	handleError("parsing templates", err)

	authRepo := repo.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo)

	indexHandler := handler.NewIndexHandler(authService, tmpl)
	expenseHandler := handler.ExpensesHandler(tmpl)
	hubHandler := handler.HubHandler(tmpl)

	m := http.NewServeMux()
	m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	m.HandleFunc("/", indexHandler.Index)
	m.Handle("GET /expenses", expenseHandler)
	m.Handle("GET /hub", hubHandler)

	s := http.Server{
		Addr:    ":1337",
		Handler: m,
	}

	err = s.ListenAndServe()
	handleError("starting server", err)
}
