package jwt

import (
	"context"
	"net/http"
	"strings"
)

type CtxKey string

// Token интерфейс для парсинга токена
type Token interface {
	Parse(string) (string, error)
}

// Auth - middleware для проверки авторизации по токену
func Auth(t Token) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header.Get("Authorization")
			if !strings.HasPrefix(h, "Bearer ") {
				http.Error(w, "no token", http.StatusUnauthorized)
				return
			}
			uid, err := t.Parse(strings.TrimPrefix(h, "Bearer "))
			if err != nil {
				http.Error(w, "bad token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), CtxKey("uid"), uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
