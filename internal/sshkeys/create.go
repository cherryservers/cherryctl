package sshkeys

import (
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		label     string
		publicKey string
	)
	sshKeyCreateCmd := &cobra.Command{
		Use:   `create --key <public_key> --label <label>`,
		Short: "Adds an SSH key for the current user's account.",
		Long:  "Adds an SSH key for the current user's account.",
		Example: `  # Adds a key labled "example-key" to the current user account.
  cherryctl ssh-key create --key ssh-rsa AAAAB3N...user@domain.com --label example-key`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.CreateSSHKey{
				Label: label,
				Key:   publicKey,
			}

			o, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not create SSH key")
			}

			header := []string{"ID", "Label", "Fingerprint", "Created"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Label, o.Fingerprint, o.Created}

			return c.Out.Output(o, header, &data)
		},
	}

	sshKeyCreateCmd.Flags().StringVarP(&label, "label", "", "", "Label of the SSH key.")
	sshKeyCreateCmd.Flags().StringVarP(&publicKey, "key", "", "", "Public SSH key string.")

	sshKeyCreateCmd.MarkFlagRequired("label")
	sshKeyCreateCmd.MarkFlagRequired("key")

	return sshKeyCreateCmd
}
