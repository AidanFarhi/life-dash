package handler

import (
	"html/template"
	"lifedash/service"
	"net/http"
)

type ExpenseHandler struct {
	t  *template.Template
	es *service.ExpenseService
}

func ExpensesHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "expenses", nil)
	}
}
