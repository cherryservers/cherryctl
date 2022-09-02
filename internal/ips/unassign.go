package ips

import (
	"fmt"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Unassign() *cobra.Command {
	var (
		ipID string
	)
	ipDetachCmd := &cobra.Command{
		Use:     `unassign ID`,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"detach", "unasign"},
		Short:   "Unassign an IP address.",
		Long:    "Unassign an IP address.",
		Example: `  # Unassign an IP address:
		cherryctl ip unassign 30c15082-a06e-4c43-bfc3-252616b46eba`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if utils.IsValidUUID(args[0]) {
				ipID = args[0]
			} else {
				fmt.Println("IP address with ID %s was not found.", args[0])
				return nil
			}

			_, err := c.Service.Unassign(ipID)
			if err != nil {
				return errors.Wrap(err, "Could not unassign IP address")
			}

			fmt.Println("IP address", ipID, "unassigned successfully.")
			return nil
		},
	}

	return ipDetachCmd
}
