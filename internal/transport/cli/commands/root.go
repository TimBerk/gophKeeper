package commands

import (
	"context"
	"database/sql"

	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/spf13/cobra"

	"github.com/TimBerk/gophKeeper/internal/platform/logger"
	"github.com/TimBerk/gophKeeper/internal/platform/migrate"
	sqliteRepo "github.com/TimBerk/gophKeeper/internal/secret/infrastructure/sqlite"
)

// RootCmd возвращает и настраивает корневую Cobra-команду.
func RootCmd(c *http.Client, sqlDB *sql.DB, repo *sqliteRepo.Repo) *cobra.Command {
	log := logger.New()

	if err := migrate.UpSQLite(sqlDB); err != nil {
		log.Fatalf("sqlite migrate: %v", err)
	}

	root := &cobra.Command{
		Use:           "gophKeeper",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// если токен есть — делаем тихий sync
			if _, err := c.LoadToken(); err == nil {
				if errSync := c.Sync(repo); errSync != nil {
					log.Warnf("sync: %v", errSync)
				}
			}
			return nil
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	root.SetContext(ctx)
	root.PersistentPostRun = func(cmd *cobra.Command, _ []string) {
		cancel()
		_ = sqlDB.Close()
	}

	root.AddCommand(cmdLogin(c, log))
	root.AddCommand(cmdAdd(c, log))
	root.AddCommand(cmdList(c, log))
	root.AddCommand(cmdSync(c, repo, log))
	root.AddCommand(cmdShow(c, repo))
	root.AddCommand(cmdVersion())

	return root
}
