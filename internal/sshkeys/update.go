package sshkeys

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Command) Update() *cobra.Command {
	var (
		label     string
		publicKey string
		id        int
	)
	sshKeyUpdateCmd := &cobra.Command{
		Use:   `update ID [--label <text>] [--key <public_key>]`,
		Args:  cobra.ExactArgs(1),
		Short: "Updates an SSH key.",
		Long:  "Updates an SSH key with either a new public key, a new label, or both.",
		Example: `  # Update SSH key:
  cherryctl ssh-key update 12345 --key AAAAB3N...user@domain.com`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			ctx := cmd.Context()

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid id: %w", err)
			}

			request := &cherrygo.UpdateSSHKey{}

			if label != "" {
				request.Label = &label
			}

			if publicKey != "" {
				request.Key = &publicKey
			}

			o, _, err := c.Client().Update(ctx, id, request)
			if err != nil {
				return errors.Wrap(err, "Could not update SSH key")
			}

			header := []string{"ID", "Label", "Fingerprint", "Created"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Label, o.Fingerprint, o.Created}

			return c.Outputer().Output(o, header, &data)
		},
	}

	sshKeyUpdateCmd.Flags().IntVarP(&id, "ssh-key-id", "i", 0, "ID of the SSH key.")
	sshKeyUpdateCmd.Flags().StringVarP(&label, "label", "", "", "Label of the SSH key.")
	sshKeyUpdateCmd.Flags().StringVarP(&publicKey, "key", "", "", "Public SSH key string.")

	sshKeyUpdateCmd.Flags().MarkDeprecated("ssh-key-id", "Pass the ID as an argument instead.")

	return sshKeyUpdateCmd
}
