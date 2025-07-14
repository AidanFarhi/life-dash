package handler

import (
	"fmt"
	"html/template"
	"lifedash/service"
	"net/http"
)

func IndexHandler(as *service.AuthService, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if user is logged in
		cookie, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("error getting cookie:", err.Error())
			t.ExecuteTemplate(w, "index", nil)
			return
		}
		validSession, err := as.ValidateSession(*cookie)
		if err != nil {
			fmt.Println("error validating session:", err.Error())
			return
		}
		if validSession {
			t.ExecuteTemplate(w, "index", nil)
			return
		}
		t.ExecuteTemplate(w, "login", nil)
	}
}