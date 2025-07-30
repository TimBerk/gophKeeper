package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/TimBerk/gophKeeper/cmd/server/app"
	logx "github.com/TimBerk/gophKeeper/internal/platform/logger"
)

func main() {
	logger := logx.New()
	a, err := app.New(logger)
	if err != nil {
		logger.Fatalf("failed to initialize application: %v", err)
	}
	if errRun := a.Run(); errRun != nil {
		logger.Fatalf("server error: %v", errRun)
	}
}
