package service

import (
	"database/sql"
	"net/http"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) ValidateSession(cookie *http.Cookie) (bool, error) {
	// handle checking the cookie using the repo
	return true, nil
}
