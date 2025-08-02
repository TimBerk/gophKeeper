package application

import (
	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"
)

// ListUseCase - UseCase для получения списка записей
type ListUseCase struct{ R sd.Repository }

// Exec - метод выполнения UseCase
func (l *ListUseCase) Exec(uid string) ([]*sd.Secret, error) { return l.R.List(uid) }
