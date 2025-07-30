package domain

import "github.com/TimBerk/gophKeeper/internal/core"

// Repository - интерфейс для работы с данными пользователя
type Repository interface {
	Save(*User) error
	ByUsername(string) (*User, error)
	ByID(core.ID) (*User, error)
}
