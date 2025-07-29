package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// New - задает новый логгер для работы с logrus
func New() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2000-01-01 00:00:00",
	})
	l.SetOutput(os.Stdout)
	return l
}
