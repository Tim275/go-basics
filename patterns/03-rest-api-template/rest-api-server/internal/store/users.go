package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Passwort string `json:"-"`
	Created  string `json:"created_at"`
}

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, password,email) VALUES($1,$2,$3) RETURNING id,
	created_at
	`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Passwort, user.Email).Scan(&user.ID, &user.Created)
	if err != nil {
		return err
	}
	return nil
}
