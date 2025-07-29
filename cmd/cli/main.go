package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimBerk/gophKeeper/internal/transport/cli/commands"
	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"

	"github.com/TimBerk/gophKeeper/internal/config"
	logx "github.com/TimBerk/gophKeeper/internal/platform/logger" // ← NEW
	"github.com/TimBerk/gophKeeper/internal/secret/infrastructure/sqlite"
)

func main() {
	_ = godotenv.Load()
	cfg := config.FromEnv()
	log := logx.New()

	db, err := sql.Open("sqlite", cfg.RootDir)
	if err != nil {
		log.Fatalf("sqlite open: %v", err)
	}
	log.Infof("sqlite path %s", cfg.RootDir)

	repo, _ := sqlite.NewWithDB(db)
	client := http.New("http://" + cfg.HTTPAddr)

	root := commands.RootCmd(client, db, repo)
	root.SilenceUsage = true
	root.SilenceErrors = true

	/* graceful-shutdown */
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cobra.OnInitialize(func() { go func() { <-ctx.Done(); log.Info("interrupted"); os.Exit(130) }() })

	log.Info("CLI started")
	if err := root.ExecuteContext(ctx); err != nil {
		log.Errorf("command error: %v", err)
		os.Exit(1)
	}
	log.Info("CLI finished")
}
