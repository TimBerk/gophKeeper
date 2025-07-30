package commands

import (
	"encoding/base64"
	"fmt"

	sd "github.com/TimBerk/gophKeeper/internal/secret/domain"
	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/spf13/cobra"
)

// cmdShow - команда для отображения записи
func cmdShow(c *http.Client, repo sd.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "show <id>",
		Short: "Показать секрет по ID с расшифровкой данных",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			sec, err := repo.GetRecord(id)
			if err != nil {
				sec, err = c.Detail(id)
				if err != nil {
					return err
				}
				_ = repo.Save(sec)
			}

			decoded, err := base64.StdEncoding.DecodeString(string(sec.Data))
			if err != nil {
				return fmt.Errorf("base64 decode: %w", err)
			}

			fmt.Printf("ID  : %s\nType: %s\nMeta: %+v\n----- DATA -----\n%s\n",
				sec.ID, sec.Type, sec.Meta, decoded)
			return nil
		},
	}
}
