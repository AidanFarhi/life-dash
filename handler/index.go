package handler

import (
	"fmt"
	"html/template"
	"lifedash/service"
	"net/http"
)

type IndexHandler struct {
	as *service.AuthService
	t  *template.Template
}

func NewIndexHandler(as *service.AuthService, t *template.Template) *IndexHandler {
	return &IndexHandler{
		as: as,
		t:  t,
	}
}

func (ih *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	// check if user is logged in
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		fmt.Println("no cookie found")
		ih.t.ExecuteTemplate(w, "login", nil)
		return
	}
	validSession, err := ih.as.ValidateSession(cookie)
	if err != nil {
		fmt.Println("error validating session:", err.Error())
		ih.t.ExecuteTemplate(w, "error", nil)
	}
	if validSession {
		fmt.Println("valid session found")
		ih.t.ExecuteTemplate(w, "index", nil)
		return
	}
	fmt.Println("no valid session found")
	ih.t.ExecuteTemplate(w, "login", nil)
}
