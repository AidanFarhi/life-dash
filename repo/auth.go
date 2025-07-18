package repo

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) SessionExists(session_id string) (bool, error) {
	sessionExists := false
	query := "SELECT EXISTS(SELECT 1 FROM session WHERE id = ?)"
	err := ar.db.QueryRow(query, session_id).Scan(&sessionExists)
	if err != nil {
		return false, err
	}
	return sessionExists, nil
}

func (ar *AuthRepo) Login(username, password string) (int, error) {
	userId := 0
	query := "SELECT id FROM user WHERE username = ? AND password = ?"
	err := ar.db.QueryRow(query, username, password).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (ar *AuthRepo) SaveSession(userId int) (string, error) {
	sessionId, err := generateSessionId()
	if err != nil {
		return "", err
	}
	query := "INSERT INTO session (id, user_id) VALUES (?, ?)"
	_, err = ar.db.Exec(query, sessionId, userId)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func generateSessionId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
