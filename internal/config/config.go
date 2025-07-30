package config

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"
)

type Config struct {
	HTTPAddr string
	JWTKey   []byte
	PGURL    string
	RootDir  string
}

// genKey — 256-битный ключ HS256/HS512.
func genKey() []byte {
	b := make([]byte, 32) // 32 байта = 256 бит
	if _, err := rand.Read(b); err != nil {
		panic("cannot generate JWT key: " + err.Error())
	}
	return b
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

	var key []byte
	if v := os.Getenv("GOVAULT_JWT"); v != "" {
		// допускаем base64 или raw-строку
		if decoded, err := base64.StdEncoding.DecodeString(v); err == nil {
			key = decoded
		} else {
			key = []byte(v)
		}
	} else {
		key = genKey()
	}

	return &Config{
		HTTPAddr: addr,
		JWTKey:   key,
		PGURL:    pg,
		RootDir:  fullPath(),
	}
}
