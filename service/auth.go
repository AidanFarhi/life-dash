package service

import (
	"lifedash/repo"
	"net/http"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) ValidateSession(cookie *http.Cookie) (bool, error) {
	sessionExists, err := as.repo.SessionExists(cookie.Value)
	if err != nil {
		return false, err
	}
	return sessionExists, nil
}
