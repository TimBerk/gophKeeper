package http

import (
	"net/http"

	"github.com/TimBerk/gophKeeper/internal/platform/jwt"
	logx "github.com/TimBerk/gophKeeper/internal/platform/logger"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// SchemaInfo интерфейс для работы с версией
type SchemaInfo interface {
	Current() (int64, error)
}

// Auth интерфейс для работы с авторизацией/регистрацией
type Auth interface {
	Register(string, string) error
	Login(string, string) (string, error)
	Parse(string) (string, error)
}

// Vault интерфейс для работы с secret-записями
type Vault interface {
	Add(string, string, []byte, map[string]string) error
	List(string) (any, error)
}

// NewRouter задаёт роутер для работы с api
func NewRouter(ai SchemaInfo, a Auth, v Vault, logger *logrus.Logger) http.Handler {
	h := NewHandler(a, v)

	r := chi.NewRouter()
	r.Use(logx.RequestLogger(logger))

	r.Get("/meta/version", h.MetaVersion(ai))
	r.Post("/register", h.Register())
	r.Post("/login", h.Login())

	r.Group(func(pr chi.Router) {
		pr.Use(jwt.Auth(a))
		pr.Post("/secret", h.AddSecret())
		pr.Get("/secret/list", h.ListSecrets())
	})
	return r
}
