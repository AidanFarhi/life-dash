package service

import "net/http"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) ValidateSession(cookie http.Cookie) (bool, error) {
	// handle checking the cookie using the repo
	return true, nil
}