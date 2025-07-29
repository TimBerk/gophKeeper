package sqlite

import (
	"database/sql"
	"encoding/json"

	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"

	_ "modernc.org/sqlite"
)

// Repo - структура репозитория для работы с запросами к таблице secret
type Repo struct{ db *sql.DB }

// NewWithDB - инициализация нового репозитория
func NewWithDB(db *sql.DB) (Repo, error) { return Repo{db: db}, nil }

// Save - сохранение записи
func (r Repo) Save(s *sd.Secret) error {
	metaJSON, _ := json.Marshal(s.Meta)
	_, err := r.db.Exec(`
		INSERT OR REPLACE INTO secrets(id,type,data,meta) VALUES (?,?,?,?)`,
		s.ID, s.Type, s.Data, string(metaJSON))
	return err
}

// List - получение списка записей
func (r Repo) List(_ string) ([]*sd.Secret, error) {
	rows, err := r.db.Query(`SELECT id,type,data,meta FROM secrets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*sd.Secret
	for rows.Next() {
		var id, typ, metaStr string
		var data []byte
		if err := rows.Scan(&id, &typ, &data, &metaStr); err != nil {
			return nil, err
		}
		var meta map[string]string
		_ = json.Unmarshal([]byte(metaStr), &meta)
		out = append(out, &sd.Secret{ID: sd.ID(id), Type: typ, Data: data, Meta: meta})
	}
	return out, rows.Err()
}
