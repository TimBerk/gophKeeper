package http

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// cacheDir - получение директории для сохранения токена
func cacheDir() string {
	if d, _ := os.UserConfigDir(); d != "" {
		return filepath.Join(d, "gophKeeper")
	}
	return filepath.Join(os.Getenv("HOME"), ".gophKeeper")
}

// tokenFile - получение пути до токена
func tokenFile() string { return filepath.Join(cacheDir(), "session.jwt") }

// ReadFileOrStdin - чтение файла из потока ввода
func ReadFileOrStdin(path string) ([]byte, error) {
	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(path)
}

// DefaultMeta - формирование базовой Meta-информации
func DefaultMeta() map[string]string {
	return map[string]string{
		"os":   runtime.GOOS,
		"arch": runtime.GOARCH,
		"ts":   time.Now().UTC().Format(time.RFC3339),
	}
}
