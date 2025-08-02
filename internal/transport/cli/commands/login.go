package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TimBerk/gophKeeper/internal/transport/cli/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// cmdLogin - реализация команды для авторизация или регистрации
func cmdLogin(c *http.Client, log *logrus.Logger) *cobra.Command {
	var user string
	var create bool

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Авторизация или регистрация (--new)",
		RunE: func(cmd *cobra.Command, args []string) error {
			reader := bufio.NewReader(os.Stdin)

			if user == "" {
				fmt.Print("username: ")
				u, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}
				user = strings.TrimSpace(u)
			}

			fmt.Print("password: ")
			pw, _ := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()

			if create {
				log.Infof("register %s", user)
				if err := c.Register(user, string(pw)); err != nil {
					return err
				}
				fmt.Println("registered OK")
			}
			log.Infof("login %s", user)
			if err := c.Login(user, string(pw)); err != nil {
				return err
			}
			fmt.Println("login OK")
			return nil
		},
	}

	cmd.Flags().StringVarP(&user, "user", "u", "", "имя пользователя")
	cmd.Flags().BoolVar(&create, "new", false, "регистрация нового пользователя")
	return cmd
}
