package sshkeys

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Command) Get() *cobra.Command {
	sshGetCmd := &cobra.Command{
		Use:   `get ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Retrieves SSH key.",
		Long:  "Retrieves the specified SSH key.",
		Example: `  # Get SSH key:
  cherryctl ssh-key get 12345`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			ctx := cmd.Context()

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid id: %w", err)
			}
			getOptions := c.GetOpts()
			if len(getOptions.Fields) == 0 {
				getOptions.Fields = []string{"ssh_key", "email"}
			}

			o, _, err := c.Client().Get(ctx, id, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not get SSH key")
			}

			header := []string{"ID", "Label", "User", "Fingerprint", "Created"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Label, o.User.Email, o.Fingerprint, o.Created}

			return c.Outputer().Output(o, header, &data)
		},
	}

	return sshGetCmd
}
