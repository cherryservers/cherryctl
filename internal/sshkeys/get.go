package sshkeys

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var sshKeyID int
	sshGetCmd := &cobra.Command{
		Use:   `get ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Retrieves ssh-key details.",
		Long:  "Retrieves the details of the specified ssh-key.",
		Example: `  # Gets the details of the specified ssh-key:
  cherryctl ssh-key get 12345`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if sshID, err := strconv.Atoi(args[0]); err == nil {
				sshKeyID = sshID
			}
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"ssh_key", "email"}
			o, _, err := c.Service.Get(sshKeyID, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not get ssh-key")
			}

			header := []string{"ID", "Label", "User", "Fingerprint", "Created"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Label, o.User.Email, o.Fingerprint, o.Created}

			return c.Out.Output(o, header, &data)
		},
	}

	return sshGetCmd
}
