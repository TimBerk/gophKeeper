package application

import (
	"github.com/TimBerk/gophKeeper/internal/core"
	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"

	"github.com/google/uuid"
)

// AddUseCase - UseCase для добавление записи
type AddUseCase struct{ R sd.Repository }

// Exec - метод выполнения UseCase
func (a *AddUseCase) Exec(uid, typ string, data []byte, meta map[string]string) error {
	return a.R.Save(&sd.Secret{
		ID: core.ID(uuid.NewString()), UserID: uid, Type: typ, Data: data, Meta: meta,
	})
}
