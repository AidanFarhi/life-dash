package handler

import (
	"html/template"
	"net/http"
)

func HomeHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "home", nil)
	}
}