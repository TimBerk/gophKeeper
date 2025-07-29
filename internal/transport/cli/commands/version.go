package commands

import (
	"fmt"

	"github.com/TimBerk/gophKeeper/pkg/version"

	"github.com/spf13/cobra"
)

// cmdVersion - команда для получения версии
func cmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Показать версию",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.String())
		},
	}
}
