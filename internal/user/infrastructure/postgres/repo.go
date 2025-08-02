package postgres

import (
	"context"
	"database/sql"

	"github.com/TimBerk/gophKeeper/internal/core"
	ud "github.com/TimBerk/gophKeeper/internal/user/domain"
)

// Repo - репозиторий для работы с данными пользователя
type Repo struct{ db *sql.DB }

// New - инициализация нового репозитория
func New(db *sql.DB) ud.Repository { return &Repo{db} }

// Save - сохранение пользователя
func (r *Repo) Save(u *ud.User) error {
	_, err := r.db.ExecContext(context.TODO(),
		`INSERT INTO users(id,username,hash) VALUES($1,$2,$3)
		 ON CONFLICT(username) DO UPDATE SET hash=EXCLUDED.hash`,
		u.ID, u.Username, u.Hash)
	return err
}

// ByUsername - поиск по логину пользователя
func (r *Repo) ByUsername(n string) (*ud.User, error) {
	var id, name string
	var h []byte
	err := r.db.QueryRowContext(context.TODO(),
		`SELECT id,username,hash FROM users WHERE username=$1`, n).
		Scan(&id, &name, &h)
	if err == sql.ErrNoRows {
		return nil, core.ErrNotFound
	}
	return &ud.User{ID: core.ID(id), Username: name, Hash: h}, err
}

// ByID - поиск по ID пользователя
func (r *Repo) ByID(id core.ID) (*ud.User, error) {
	var name string
	var h []byte
	err := r.db.QueryRowContext(context.TODO(),
		`SELECT username,hash FROM users WHERE id=$1`, id).
		Scan(&name, &h)
	if err == sql.ErrNoRows {
		return nil, core.ErrNotFound
	}
	return &ud.User{ID: id, Username: name, Hash: h}, err
}
