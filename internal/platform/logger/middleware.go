package logger

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// wrap - обработчик ответов
type wrap struct {
	http.ResponseWriter
	status int
}

// WriteHeader - добавляет код к ответу запроса
func (w *wrap) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// RequestLogger - middleware для логирования запросов к сервису
func RequestLogger(l *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &wrap{ResponseWriter: w, status: 200}
			next.ServeHTTP(ww, r)
			l.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"ip":     r.RemoteAddr,
				"status": ww.status,
				"lat":    time.Since(start).Milliseconds(),
			}).Info("request")
		})
	}
}
