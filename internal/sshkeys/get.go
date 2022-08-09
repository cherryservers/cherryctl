package sshkeys

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var sshKeyID int
	sshGetCmd := &cobra.Command{
		Use:   `get [-i <ssh_key_id>]`,
		Short: "Retrieves ssh-key details.",
		Long:  "Retrieves the details of the specified ssh-key.",
		Example: `  # Gets the details of the specified ssh-key:
  cherryctl ssh-key get -i 12345`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
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

	sshGetCmd.Flags().IntVarP(&sshKeyID, "ssh-key-id", "i", 0, "The ID of ssh-key.")
	_ = sshGetCmd.MarkFlagRequired("ssh-key-id")

	return sshGetCmd
}
