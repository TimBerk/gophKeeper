package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/TimBerk/gophKeeper/internal/config"
	web "github.com/TimBerk/gophKeeper/internal/platform/http"
	"github.com/TimBerk/gophKeeper/internal/platform/jwt"
	"github.com/TimBerk/gophKeeper/internal/platform/migrate"
	schemaPG "github.com/TimBerk/gophKeeper/internal/schema/infrastructure/postgres"
	secApp "github.com/TimBerk/gophKeeper/internal/secret/application"
	secPG "github.com/TimBerk/gophKeeper/internal/secret/infrastructure/postgres"
	userApp "github.com/TimBerk/gophKeeper/internal/user/application"
	userPG "github.com/TimBerk/gophKeeper/internal/user/infrastructure/postgres"

	"github.com/joho/godotenv"
)

type App struct {
	cfg    *config.Config
	logger *logrus.Logger
	db     *sql.DB
	srv    *http.Server
}

// New читает конфиг, прокидывает зависимости, мигрирует базу, собирает роутер и http.Server.
func New(logger *logrus.Logger) (*App, error) {
	_ = godotenv.Load()
	cfg := config.FromEnv()

	db, err := sql.Open("pgx", cfg.PGURL)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(5 * time.Minute)

	if errUpPostgres := migrate.UpPostgres(db, "migrations/server"); errUpPostgres != nil {
		return nil, errUpPostgres
	}

	// 6. UseCases и репозитории
	authUC := userApp.Auth{
		Repo: userPG.New(db),
		JWT:  jwt.New(cfg.JWTKey),
	}
	vaultRepo := secPG.New(db)
	addUC := secApp.AddUseCase{R: vaultRepo}
	listUC := secApp.ListUseCase{R: vaultRepo}
	detailUC := secApp.DetailUseCase{R: vaultRepo}
	vaultAdapter := secApp.VaultAdapter{AddUC: &addUC, ListUC: &listUC, DetailUC: &detailUC}
	schemaInfo := schemaPG.New(db)

	router := web.NewRouter(schemaInfo, &authUC, &vaultAdapter, logger)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &App{
		cfg:    cfg,
		logger: logger,
		db:     db,
		srv:    srv,
	}, nil
}

// Run запускает HTTP-сервер и слушает сигналы остановки
func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		a.logger.Printf("Shutting down server…")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := a.srv.Shutdown(shutdownCtx); err != nil {
			a.logger.Printf("Error during shutdown: %v", err)
		}
	}()

	a.logger.Printf("HTTP listening on %s", a.cfg.HTTPAddr)
	err := a.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
