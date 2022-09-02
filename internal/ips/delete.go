package ips

import (
	"fmt"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var ipID string
	var force bool
	deleteIpCmd := &cobra.Command{
		Use:   `delete ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Delete an IP address.",
		Long:  "Deletes the specified IP address with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified IP:
  cherryctl ip delete 30c15082-a06e-4c43-bfc3-252616b46eba
  >
  ✔ Are you sure you want to delete IP address 30c15082-a06e-4c43-bfc3-252616b46eba: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl ip delete 30c15082-a06e-4c43-bfc3-252616b46eba -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if utils.IsValidUUID(args[0]) {
				ipID = args[0]
			} else {
				fmt.Println("IP address with ID %s was not found.", args[0])
				return nil
			}

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

	deleteIpCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the server deletion.")

	return deleteIpCmd
}
