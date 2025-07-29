package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

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
	defer rows.Close()

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
		out = append(out, &sd.Secret{ID: sd.ID(id), UserID: uid, Type: typ, Data: data, Meta: meta})
	}
	return out, rows.Err()
}
