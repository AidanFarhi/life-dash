package handler

import (
	"html/template"
	"net/http"
)

type LoginHandler struct {
	t *template.Template
}

func NewLoginHandler(t *template.Template) *LoginHandler {
	return &LoginHandler{t}
}

func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	lh.t.ExecuteTemplate(w, "login", nil)
}
