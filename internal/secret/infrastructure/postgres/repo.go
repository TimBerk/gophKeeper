package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/TimBerk/gophKeeper/internal/core"
	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"
)

// Repo - структура репозитория для работы с запросами к таблице secret
type Repo struct{ db *sql.DB }

// New - инициализация нового репозитория
func New(db *sql.DB) sd.Repository { return &Repo{db} }

// Save - сохранение записи
func (r *Repo) Save(s *sd.Secret) error {
	_, err := r.db.ExecContext(context.TODO(),
		`INSERT INTO secrets(id,user_id,type,data,meta)
		 VALUES($1,$2,$3,$4,$5)
		 ON CONFLICT(id) DO UPDATE SET type=EXCLUDED.type,data=EXCLUDED.data,meta=EXCLUDED.meta`,
		s.ID, s.UserID, s.Type, s.Data, s.Meta)
	return err
}

// List - получение списка записей
func (r *Repo) List(uid string) ([]*sd.Secret, error) {
	rows, err := r.db.QueryContext(context.TODO(),
		`SELECT id,type,data,meta FROM secrets WHERE user_id=$1`, uid)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Errorf("rows close: %v", errClose)
		}
	}()

	var out []*sd.Secret
	for rows.Next() {
		var id, typ string
		var data []byte
		var metaBytes []byte
		if errScan := rows.Scan(&id, &typ, &data, &metaBytes); errScan != nil {
			return nil, errScan
		}
		var meta map[string]string
		if len(metaBytes) > 0 {
			if err := json.Unmarshal(metaBytes, &meta); err != nil {
				return nil, err
			}
		}
		out = append(out, &sd.Secret{ID: core.ID(id), UserID: uid, Type: typ, Data: data, Meta: meta})
	}

	if errRows := rows.Err(); errRows != nil {
		return nil, fmt.Errorf("row iteration: %w", errRows)
	}
	return out, nil
}

// GetRecord - получение записи по ID
func (r *Repo) GetRecord(id string) (*sd.Secret, error) {
	row := r.db.QueryRow(`SELECT user_id,type,data,meta FROM secrets WHERE id=$1`, id)
	var uid, typ string
	var data []byte
	var meta map[string]string
	if err := row.Scan(&uid, &typ, &data, &meta); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrNotFound
		}
		return nil, err
	}
	return &sd.Secret{ID: core.ID(id), UserID: uid, Type: typ, Data: data, Meta: meta}, nil
}
