package domain

import "github.com/TimBerk/gophKeeper/internal/core"

// Secret сущность для работы с secret-записями
type Secret struct {
	ID     core.ID
	UserID string
	Type   string
	Data   []byte
	Meta   map[string]string
}
