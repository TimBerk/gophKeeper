package migrate

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

// UpPostgres применяет все недостающие миграции.
// Возвращает ошибку, если база «позади» или недоступна.
func UpPostgres(db *sql.DB, dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	return goose.Up(db, dir)
}
