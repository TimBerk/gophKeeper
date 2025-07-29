package commands

import (
	"encoding/base64"

	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cmdAdd - реализация команды добавления записи
func cmdAdd(c *http.Client, log *logrus.Logger) *cobra.Command {
	var typ, file string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Добавить секрет",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := http.ReadFileOrStdin(file)
			if err != nil {
				return err
			}
			log.Infof("add %s (%d bytes)", typ, len(data))
			b64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
			base64.StdEncoding.Encode(b64, data)
			return c.Add(typ, b64, http.DefaultMeta())
		},
	}
	cmd.Flags().StringVar(&typ, "type", "text", "тип данных")
	cmd.Flags().StringVarP(&file, "file", "f", "-", "файл или - для stdin")
	return cmd
}
