package service

import (
	"database/sql"
	"fmt"
	"lifedash/repo"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) ValidateSession(sessionId string) (bool, error) {
	fmt.Println(sessionId)
	sessionExists, err := as.repo.SessionExists(sessionId)
	fmt.Println("auth service:", sessionId, err)
	if err != nil {
		return false, err
	}
	return sessionExists, nil
}

func (as *AuthService) Login(username, password string) (bool, string, error) {
	userId, err := as.repo.Login(username, password)
	if err == sql.ErrNoRows {
		return false, "", nil
	}
	if err != nil {
		return false, "", err
	}
	sessionId, err := as.repo.SaveSession(userId)
	if err != nil {
		return false, "", err
	}
	return true, sessionId, nil
}
