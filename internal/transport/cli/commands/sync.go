package commands

import (
	"fmt"

	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"
	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cmdSync - команда для синхронизации
func cmdSync(c *http.Client, repo sd.Repository, log *logrus.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Принудительная синхронизация данных c сервером",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("start sync with server")
			if err := c.Sync(repo); err != nil {
				log.WithError(err).Error("sync failed")
				return err
			}
			fmt.Println("Синхронизация завершена успешно")
			log.Info("sync completed")
			return nil
		},
	}
}
