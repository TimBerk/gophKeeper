package migrate

import (
	"database/sql"

	"github.com/TimBerk/gophKeeper/migrations/local"

	"github.com/pressly/goose/v3"
)

// UpSQLite применяет встроенные миграции к локальной БД.
func UpSQLite(db *sql.DB) error {
	goose.SetBaseFS(local.FS)
	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}
	return goose.Up(db, ".")
}
