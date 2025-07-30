package application

import (
	"time"

	"github.com/TimBerk/gophKeeper/internal/core"
	ujwt "github.com/TimBerk/gophKeeper/internal/platform/jwt"
	ud "github.com/TimBerk/gophKeeper/internal/user/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Auth - структура для работы с авторизацией/регистрацией
type Auth struct {
	Repo ud.Repository
	JWT  *ujwt.Service
}

// Register - реализация поиска и добавления пользователя при регистрации
func (a *Auth) Register(name, pw string) error {
	if _, err := a.Repo.ByUsername(name); err == nil {
		return core.ErrExists
	}
	u, err := ud.New(core.ID(uuid.NewString()), name, pw)
	if err != nil {
		return err
	}
	return a.Repo.Save(u)
}

// Login - реализация поиска и проверки пользователя при авторизации
func (a *Auth) Login(name, pw string) (string, error) {
	u, err := a.Repo.ByUsername(name)
	if err != nil {
		return "", core.ErrNotFound
	}
	if !u.Check(pw) {
		return "", core.ErrBadPassword
	}
	return a.JWT.Sign(jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
}

func (a *Auth) Parse(tok string) (string, error) {
	return a.JWT.ParseSub(tok)
}
