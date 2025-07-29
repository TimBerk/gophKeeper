package commands

import (
	"encoding/json"
	"fmt"

	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cmdList - реализация команды получения списка
func cmdList(c *http.Client, log *logrus.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Список секретов",
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := c.List()
			if err != nil {
				return err
			}
			out, _ := json.MarshalIndent(list, "", "  ")
			log.Infof("list returned %d items", len(list))
			fmt.Println(string(out))
			return nil
		},
	}
}
