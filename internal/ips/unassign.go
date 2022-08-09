package ips

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Unassign() *cobra.Command {
	var (
		ipID string
	)
	ipDetachCmd := &cobra.Command{
		Use:     `unassign -i <ip_address_id>`,
		Aliases: []string{"detach", "unasign"},
		Short:   "Unassign an IP address.",
		Long:    "Unassign an IP address.",
		Example: `  # Unassign an IP address:
		cherryctl ip unassign -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			_, err := c.Service.Unassign(ipID)
			if err != nil {
				return errors.Wrap(err, "Could not unassign IP address")
			}

			fmt.Println("IP address", ipID, "unassigned successfully.")
			return nil
		},
	}

	ipDetachCmd.Flags().StringVarP(&ipID, "ip-address-id", "i", "", "The ID of an IP address.")

	ipDetachCmd.MarkFlagRequired("ip-address-id")

	return ipDetachCmd
}
