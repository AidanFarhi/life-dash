package service

import (
	"lifedash/repo"
	"net/http"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) ValidateSession(cookie *http.Cookie) (bool, error) {
	// handle checking the cookie using the repo
	return true, nil
}
