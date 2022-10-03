package storages

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var storageID int
	var force bool
	deleteStorageCmd := &cobra.Command{
		Use:   `delete ID [f]`,
		Args:  cobra.ExactArgs(1),
		Short: "Delete a storage.",
		Long:  "Deletes the specified storage with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified storage:
  cherryctl storage delete 12345
  >
  ✔ Are you sure you want to delete storage 12345: y
  		
  # Deletes a storage, skipping confirmation:
  cherryctl storage delete 12345 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if storID, err := strconv.Atoi(args[0]); err == nil {
				storageID = storID
			}
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete storage %d? ", storageID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}

			_, err := c.Service.Delete(storageID)
			if err != nil {
				return errors.Wrap(err, "Could not delete storage")
			}

			fmt.Println("Storage", storageID, "successfully deleted.")
			return nil
		},
	}

	deleteStorageCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the storage deletion.")

	return deleteStorageCmd
}
