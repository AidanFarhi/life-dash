package handler

import (
	"html/template"
	"lifedash/service"
	"net/http"
)

type LoginHandler struct {
	t  *template.Template
	as *service.AuthService
}

func NewLoginHandler(t *template.Template, as *service.AuthService) *LoginHandler {
	return &LoginHandler{
		t:  t,
		as: as,
	}
}

func (lh *LoginHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	lh.t.ExecuteTemplate(w, "login", nil)
}

func (lh *LoginHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	loginSuccessful, sessionId, err := lh.as.Login(username, password)
	if err != nil {
		// TODO: custom error page
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loginSuccessful {
		lh.t.ExecuteTemplate(w, "login", nil)
		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		// Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
