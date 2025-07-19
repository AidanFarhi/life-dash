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
	alreadyLoggedIn, err := lh.AlreadyLoggedIn(w, r)
	if err != nil {
		// TODO: show an error page here
		lh.t.ExecuteTemplate(w, "login", nil)
		return
	}
	if alreadyLoggedIn {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	lh.t.ExecuteTemplate(w, "login", nil)
}

func (lh *LoginHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	alreadyLoggedIn, err := lh.AlreadyLoggedIn(w, r)
	if err != nil {
		// TODO: show an error page here
		lh.t.ExecuteTemplate(w, "login", nil)
		return
	}
	if alreadyLoggedIn {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	loginSuccessful, sessionId, err := lh.as.Login(username, password)
	if err != nil {
		// TODO: show an error page here
		lh.t.ExecuteTemplate(w, "login", nil)
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

func (lh *LoginHandler) AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) (bool, error) {
	cookie, err := r.Cookie("session_id")
	if err != http.ErrNoCookie {
		sessionExists, err := lh.as.ValidateSession(cookie.Value)
		if err != nil {
			return false, err
		}
		if sessionExists {
			return true, nil
		}
	}
	return false, nil
}
