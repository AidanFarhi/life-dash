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
	// TODO: implement me
	//sessionExists, err := as.repo.SessionExists(cookie.Value)
	return true, nil
}
