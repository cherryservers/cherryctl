package ips

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var ipID string
	var force bool
	deleteIpCmd := &cobra.Command{
		Use:   `delete -i <ip_address_id>`,
		Short: "Delete an IP address.",
		Long:  "Deletes the specified IP address with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified IP:
  cherryctl ip delete -i 30c15082-a06e-4c43-bfc3-252616b46eba
  >
  ✔ Are you sure you want to delete IP address 30c15082-a06e-4c43-bfc3-252616b46eba: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl ip delete -f -i 30c15082-a06e-4c43-bfc3-252616b46eba`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete IP address %s: ", ipID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			_, err := c.Service.Remove(ipID)
			if err != nil {
				return errors.Wrap(err, "Could not delete IP address")
			}

			fmt.Println("IP address", ipID, "successfully deleted.")
			return nil
		},
	}

	deleteIpCmd.Flags().StringVarP(&ipID, "ip-address-id", "i", "", "The ID of a IP address.")
	deleteIpCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the server deletion.")

	_ = deleteIpCmd.MarkFlagRequired("ip-address-id")

	return deleteIpCmd
}
