package service

import (
	"database/sql"
	"errors"
	"lifedash/repo"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) ValidateSession(sessionId string) (int, error) {
	sessionExists, err := as.repo.GetUserIdFromSession(sessionId)
	if err != nil {
		return 0, err
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

func (as *AuthService) Logout(sessionId string) error {
	userId, err := as.ValidateSession(sessionId)
	if err != nil {
		return err
	}
	if userId == 0 {
		return errors.New("session not found")
	}
	err = as.repo.DeleteSession(sessionId)
	if err != nil {
		return err
	}
	return nil
}
