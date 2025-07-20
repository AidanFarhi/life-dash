package handler

import (
	"fmt"
	"html/template"
	"lifedash/service"
	"net/http"
)

type AuthenticationHandler struct {
	t  *template.Template
	as *service.AuthService
}

func NewAuthenticationHandler(t *template.Template, as *service.AuthService) *AuthenticationHandler {
	return &AuthenticationHandler{
		t:  t,
		as: as,
	}
}

func (lh *AuthenticationHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	lh.t.ExecuteTemplate(w, "login", nil)
}

func (lh *AuthenticationHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
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

func (lh *AuthenticationHandler) PostLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		if err := lh.as.Logout(cookie.Value); err != nil {
			fmt.Println("logout error:", err)
		}
		expiredCookie := &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, expiredCookie)
	}
	w.Header().Set("HX-Redirect", "/login")
}
