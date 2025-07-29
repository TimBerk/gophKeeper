package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	HTTPAddr string
	JWTKey   []byte
	PGURL    string
	RootDir  string
}

// projectRoot возвращает:
//   - значение переменной RootDir, если она задана ;
//   - иначе – текущий рабочий каталог (from os.Getwd()).
func projectRoot() string {
	if dir := os.Getenv("RootDir"); dir != "" {
		return dir
	}
	wd, err := os.Getwd()
	if err != nil {
		return "." // крайний случай – используем точку
	}
	return wd
}

// fullPath возвращает полный путь к файлу с локальной БД.
func fullPath() string {
	return filepath.Join(projectRoot(), "data.db")
}

// FromEnv - получение переменных окружения
func FromEnv() *Config {
	addr := ":8080"
	if v := os.Getenv("GOVAULT_HTTP"); v != "" {
		addr = v
	}

	pg := "postgres://keeper:keeper@localhost/keeper?sslmode=disable"
	if v := os.Getenv("GOVAULT_PG"); v != "" {
		pg = v
	}

	return &Config{
		HTTPAddr: addr,
		JWTKey:   []byte("change-me"),
		PGURL:    pg,
		RootDir:  fullPath(),
	}
}
