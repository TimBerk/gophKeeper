package domain

// ID сущность для работы с ID-записи
type ID string

// Secret сущность для работы с secret-записями
type Secret struct {
	ID     ID
	UserID string
	Type   string
	Data   []byte
	Meta   map[string]string
}
