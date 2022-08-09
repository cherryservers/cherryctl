package sshkeys

import (
	"strconv"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		sshKeyID  int
		label     string
		publicKey string
	)
	sshKeyUpdateCmd := &cobra.Command{
		Use:   `update -i <ssh_key_id> [--label] [--key <public_key>]`,
		Short: "Updates an SSH key.",
		Long:  "Updates an SSH key with either a new public key, a new label, or both.",
		Example: `  # Update team to change currency to EUR:
  cherryctl ssh-key update -i 12345 --key AAAAB3N...user@domain.com`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.UpdateSSHKey{}

			if label != "" {
				request.Label = &label
			}

			if publicKey != "" {
				request.Key = &publicKey
			}

			o, _, err := c.Service.Update(sshKeyID, request)
			if err != nil {
				return errors.Wrap(err, "Could not update SSH key")
			}

			header := []string{"ID", "Label", "Fingerprint", "Created"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Label, o.Fingerprint, o.Created}

			return c.Out.Output(o, header, &data)
		},
	}

	sshKeyUpdateCmd.Flags().IntVarP(&sshKeyID, "ssh-key-id", "i", 0, "ID of the SSH key.")
	sshKeyUpdateCmd.Flags().StringVarP(&label, "label", "", "", "Label of the SSH key.")
	sshKeyUpdateCmd.Flags().StringVarP(&publicKey, "key", "", "", "Public SSH key string.")

	sshKeyUpdateCmd.MarkFlagRequired("ssh-key-id")

	return sshKeyUpdateCmd
}
