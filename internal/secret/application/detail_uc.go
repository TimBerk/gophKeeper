package application

import (
	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"
)

// DetailUseCase - UseCase для получения записи
type DetailUseCase struct{ R sd.Repository }

// Exec - метод выполнения UseCase
func (g *DetailUseCase) Exec(uid string) (*sd.Secret, error) { return g.R.GetRecord(uid) }
