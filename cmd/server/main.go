package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	logx "github.com/TimBerk/gophKeeper/internal/platform/logger"
	"github.com/TimBerk/gophKeeper/internal/platform/migrate"

	"github.com/joho/godotenv"

	"github.com/TimBerk/gophKeeper/internal/config"
	web "github.com/TimBerk/gophKeeper/internal/platform/http"
	"github.com/TimBerk/gophKeeper/internal/platform/jwt"
	schemaPG "github.com/TimBerk/gophKeeper/internal/schema/infrastructure/postgres"
	secApp "github.com/TimBerk/gophKeeper/internal/secret/application"
	secPG "github.com/TimBerk/gophKeeper/internal/secret/infrastructure/postgres"
	userApp "github.com/TimBerk/gophKeeper/internal/user/application"
	userPG "github.com/TimBerk/gophKeeper/internal/user/infrastructure/postgres"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	_ = godotenv.Load()
	cfg := config.FromEnv()
	log := logx.New()

	db, err := sql.Open("pgx", cfg.PGURL)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(5 * time.Minute)
	if errMigrate := migrate.UpPostgres(db, "migrations/server"); errMigrate != nil {
		log.Fatal(errMigrate)
	}

	authUC := userApp.Auth{Repo: userPG.New(db), JWT: jwt.New(cfg.JWTKey)}
	vaultRepo := secPG.New(db)
	addUC := secApp.AddUseCase{R: vaultRepo}
	listUC := secApp.ListUseCase{R: vaultRepo}
	vault := secApp.VaultAdapter{AddUC: &addUC, ListUC: &listUC}
	schemaInfo := schemaPG.New(db)

	router := web.NewRouter(schemaInfo, &authUC, &vault, log)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(shCtx)
	}()

	log.Printf("HTTP listening on %s", cfg.HTTPAddr)
	if errServe := srv.ListenAndServe(); errServe != nil && !errors.Is(errServe, http.ErrServerClosed) {
		log.Fatal(errServe)
	}
}
