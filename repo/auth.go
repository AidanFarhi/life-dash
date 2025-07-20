package repo

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) SessionExists(sessionId string) (bool, error) {
	sessionExists := false
	query := "SELECT EXISTS(SELECT 1 FROM session WHERE id = ?)"
	err := ar.db.QueryRow(query, sessionId).Scan(&sessionExists)
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

func (ar *AuthRepo) DeleteSession(sessionId string) error {
	query := "DELETE FROM session WHERE id = ?"
	_, err := ar.db.Exec(query, sessionId)
	fmt.Println(err)
	return err
}

func generateSessionId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
