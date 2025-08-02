package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// Service структура для работы с токеном
type Service struct{ key []byte }

// New создает новый обработчик для работы с токеном
func New(key []byte) *Service { return &Service{key} }

// Sign задает токен
func (s *Service) Sign(claims jwt.MapClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.key)
}

// ParseSub парсит токен
func (s *Service) ParseSub(tok string) (string, error) {
	t, err := jwt.Parse(tok, func(_ *jwt.Token) (any, error) { return s.key, nil })
	if err != nil || !t.Valid {
		return "", errors.New("invalid token")
	}
	sub, _ := t.Claims.(jwt.MapClaims)["sub"].(string)
	return sub, nil
}
