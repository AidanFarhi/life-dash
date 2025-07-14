package handler

import (
	"html/template"
	"net/http"
)

func ExpensesHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "expenses", nil)
	}
}
