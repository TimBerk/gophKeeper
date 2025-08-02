package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// New - задает новый логгер для работы с logrus
func New() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	return l
}

// Close - логирование ошибки закрытия
func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		logrus.Warnf("close: %v", err)
	}
}
