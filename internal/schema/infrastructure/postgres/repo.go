package postgres

import (
	"database/sql"

	"github.com/pressly/goose/v3"

	sd "github.com/TimBerk/gophKeeper/internal/schema/domain"
)

// SchemaInfo структура для работы с версией
type SchemaInfo struct {
	db *sql.DB
}

// New создает обработчик для версий
func New(db *sql.DB) sd.Repository { return &SchemaInfo{db} }

// Current получение версии из БД
func (schema *SchemaInfo) Current() (int64, error) { return goose.EnsureDBVersion(schema.db) }
