package handler

import (
	"html/template"
	"net/http"
)

type IndexHandler struct {
	t *template.Template
}

func NewIndexHandler(t *template.Template) *IndexHandler {
	return &IndexHandler{t}
}

func (ih *IndexHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	ih.t.ExecuteTemplate(w, "index", nil)
}
