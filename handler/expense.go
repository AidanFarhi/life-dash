package handler

import (
	"encoding/json"
	"html/template"
	"lifedash/service"
	"net/http"
)

const userIdKey = "userId"

type ExpenseHandler struct {
	t  *template.Template
	es *service.ExpenseService
}

func NewExpenseHandler(t *template.Template, es *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		t:  t,
		es: es,
	}
}

func (eh *ExpenseHandler) GetExpensesJSON(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIdKey).(int)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	expenses, err := eh.es.GetExpensesForUser(userId)
	if err != nil {
		http.Error(w, "error getting expenses", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		http.Error(w, "failed to encode expenses", http.StatusInternalServerError)
	}
}

func (eh *ExpenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	eh.t.ExecuteTemplate(w, "expenses", nil)
}
