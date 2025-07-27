package middleware

import (
	"context"
	"lifedash/service"
	"net/http"
)

// TODO: create an actual logger

const userIdKey = "userId"

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
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		userId, err := am.as.ValidateSession(cookie.Value)
		if err != nil {
			// TODO: custom error page
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), userIdKey, userId)
		next(w, r.WithContext(ctx))
	}
}

func (am *AuthMiddleware) RedirectIfLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == nil {
			userId, err := am.as.ValidateSession(cookie.Value)
			if err != nil {
				// TODO: custom error page
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if userId == 0 {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		next(w, r)
	}
}
