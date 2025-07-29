# GophKeeper

<p align="left">
    <a href="https://go.dev/" target="blank">
        <img src="https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge" />
    </a>
    <a href="https://www.docker.com/" target="blank">
        <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" />
    </a>
    <a href="https://www.postgresql.org/" target="blank">
        <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white"/>
    </a>
</p>

Простая реализация сервера и консольного приложения для возможности хранить секретные данные

## Built With

* [Logrus](https://github.com/sirupsen/logrus) - пакет для логирования.
* [Chi](https://github.com/go-chi/chi) — лёгкий и идиоматичный HTTP-роутер для построения REST-сервисов в Go, поддерживающий middlewares и сложную маршрутизацию.
* [Golang-JWT/jwt](https://github.com/golang-jwt/jwt) — популярная библиотека для генерации и верификации JSON Web Token (JWT) в Go, поддерживает HMAC, RSA и другие алгоритмы.
* [Google UUID](https://pkg.go.dev/github.com/google/uuid) — пакет для генерации и парсинга UUID v1, v4 и других, реализует RFC 4122.
* [PGX](https://github.com/jackc/pgx) — высокопроизводительный драйвер PostgreSQL для Go; поддерживает как низкоуровневый API, так и стандартный database/sql.
* [Godotenv](https://github.com/joho/godotenv) — переносит env-переменные из .env-файлов в окружение процесса.
* [Goose](https://github.com/pressly/goose) — инструмент и библиотека для управления миграциями базы данных и версионирования схемы .
* [Cobra](https://github.com/spf13/cobra) — наиболее популярный фреймворк для построения многоуровневых CLI-приложений в Go (поддержка вложенных команд, автодокументация, bash/zsh completion).
* [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) — набор современных и устоявшихся криптопримитивов и реализаций криптографических API для Go.
* [golang.org/x/term](https://pkg.go.dev/golang.org/x/term) — библиотека для работы с терминальными устройствами (например, безопасный ввод паролей).
* [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) — чисто-Go-драйвер для работы с SQLite, не требует CGO/компилятора C (в отличие от go-sqlite3).


## Make команды

* **lint** - запуск линтера.
* **format** - форматирование импортов.
* **run** - запуск сервера.
* **dbu** - запуск контейнера с БД.
* **dbd** - остановка контейнера с БД.

## Пример работы с консольным приложением

```bash
# 1. Регистрация (login --new)
go run ./cmd/cli login --new -u alice

# 2. Аторизация (login)
go run ./cmd/cli login -u alice

# 3. Добавление секрета (add)
go run ./cmd/cli add --type text -f README.md

# 4. Получить список всех секретов (list)
go run ./cmd/cli list

# 5. Принудительная синхронизация (sync)
go run ./cmd/cli sync

# 6. Показать информацию о версии клиента (version)
go run ./cmd/cli version
```

```env
GOVAULT_HTTP=localhost:8080
GOVAULT_PG=postgres://keeper:keeper@localhost:5432/keeper?sslmode=disable
GOVAULT_JWT=super-secret-key
```