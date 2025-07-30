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

func (eh *ExpenseHandler) GetAggregatedExpensesJSON(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIdKey).(int)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	aggregatedExpenses, err := eh.es.GetAggregatedExpensesForUser(userId)
	if err != nil {
		http.Error(w, "error getting aggregatedExpenses", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(aggregatedExpenses); err != nil {
		http.Error(w, "failed to encode aggregatedExpenses", http.StatusInternalServerError)
	}
}

func (eh *ExpenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(userIdKey).(int)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	allExpenses, err := eh.es.GetAllExpensesForUser(userId)
	if err != nil {
		http.Error(w, "error getting allExpenses", http.StatusInternalServerError)
		return
	}
	eh.t.ExecuteTemplate(w, "expenses", allExpenses)
}
