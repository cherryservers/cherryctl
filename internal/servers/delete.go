package servers

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var serverID int
	var force bool
	deleteServerCmd := &cobra.Command{
		Use:   `delete -i <server_id>`,
		Short: "Delete a server.",
		Long:  "Deletes the specified server with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified server:
  cherryctl server delete -i 12345
  >
  ✔ Are you sure you want to delete server 12345: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl server delete -f -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete server %d: ", serverID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			_, _, err := c.Service.Delete(serverID)
			if err != nil {
				return errors.Wrap(err, "Could not delete Server")
			}

			fmt.Println("Server", serverID, "successfully deleted.")
			return nil
		},
	}

	deleteServerCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	deleteServerCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the server deletion.")

	_ = deleteServerCmd.MarkFlagRequired("server-id")

	return deleteServerCmd
}
