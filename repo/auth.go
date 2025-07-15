package repo

import "database/sql"

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
	query := "SELECT EXISTS(SELECT 1 FROM sessions WHERE session_id = ?)"
	err := ar.db.QueryRow(query, session_id).Scan(&sessionExists)
	if err != nil {
		return false, err
	}
	return sessionExists, nil
}
