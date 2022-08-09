package storages

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var storageID int
	var force bool
	deleteStorageCmd := &cobra.Command{
		Use:   `delete -i <storage_id>`,
		Short: "Delete a storage.",
		Long:  "Deletes the specified storage with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified storage:
  cherryctl storage delete -i 12345
  >
  ✔ Are you sure you want to delete storage 12345: y
  		
  # Deletes a storage, skipping confirmation:
  cherryctl storage delete -f -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
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

	deleteStorageCmd.Flags().IntVarP(&storageID, "storage-id", "i", 0, "The ID of a storage volume.")
	deleteStorageCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the storage deletion.")

	_ = deleteStorageCmd.MarkFlagRequired("storage-id")

	return deleteStorageCmd
}
