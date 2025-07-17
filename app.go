package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"lifedash/handler"
	"lifedash/middleware"
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

	// init repos
	authRepo := repo.NewAuthRepo(db)

	// init services
	authService := service.NewAuthService(authRepo)

	// init middleware
	authMiddlWare := middleware.NewAuthMiddleware(authService)

	// init handlers
	indexHandler := handler.NewIndexHandler(tmpl)
	loginHandler := handler.NewLoginHandler(tmpl)
	expenseHandler := handler.ExpensesHandler(tmpl)
	hubHandler := handler.HubHandler(tmpl)

	// create multiplexer
	mux := http.NewServeMux()

	// register handlers and apply middleware
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/", authMiddlWare.RequireAuth(http.HandlerFunc(indexHandler.GetIndex)))
	mux.HandleFunc("GET /login", loginHandler.GetLogin)
	mux.Handle("GET /expenses", expenseHandler)
	mux.Handle("GET /hub", hubHandler)

	// config server
	s := http.Server{
		Addr:    ":1337",
		Handler: mux,
	}

	// start server
	err = s.ListenAndServe()
	handleError("starting server", err)
}
