package handler

import (
	"html/template"
	"net/http"
)

func HubHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "hub", nil)
	}
}