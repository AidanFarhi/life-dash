package middleware

import (
	"fmt"
	"lifedash/service"
	"net/http"
)

// TODO: create an actual logger

type AuthMiddleware struct {
	as *service.AuthService
}

func NewAuthMiddleware(as *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{as}
}

func (am *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println("no session_id cookie found")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		sessionValid, err := am.as.ValidateSession(cookie.Value)
		if err != nil {
			// TODO: custom error page
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !sessionValid {
			fmt.Println("session not valid")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}

func (am *AuthMiddleware) RedirectIfLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == nil {
			sessionExists, err := am.as.ValidateSession(cookie.Value)
			if err != nil {
				// TODO: custom error page
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if sessionExists {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		next(w, r)
	}
}
