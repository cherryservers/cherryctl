package servers

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var force bool
	deleteServerCmd := &cobra.Command{
		Use:   `delete ID [-f]`,
		Args:  cobra.ExactArgs(1),
		Short: "Delete a server.",
		Long:  "Deletes the specified server with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified server:
  cherryctl server delete 12345
  >
  ✔ Are you sure you want to delete server 12345: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl server delete 12345 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete server %s: ", args[0]),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.Delete(serverID)
				if err != nil {
					return errors.Wrap(err, "Could not delete Server")
				}

				fmt.Println("Server", serverID, "successfully deleted.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}

	deleteServerCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the server deletion.")

	return deleteServerCmd
}
