package sshkeys

import (
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
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
		Use:   `update ID [--label <text>] [--key <public_key>]`,
		Args:  cobra.ExactArgs(1),
		Short: "Updates an SSH key.",
		Long:  "Updates an SSH key with either a new public key, a new label, or both.",
		Example: `  # Update team to change currency to EUR:
  cherryctl ssh-key update 12345 --key AAAAB3N...user@domain.com`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if sshID, err := strconv.Atoi(args[0]); err == nil {
				sshKeyID = sshID
			}
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
