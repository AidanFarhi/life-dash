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

func (am *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println("no session_id cookie found")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		sessionValid, err := am.as.ValidateSession(cookie)
		if err != nil {
			fmt.Println("error validating session")
			return
		}
		if !sessionValid {
			fmt.Println("session not valid")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
