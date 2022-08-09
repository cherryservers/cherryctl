package storages

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Detach() *cobra.Command {
	var (
		storageID int
	)
	storageDetachCmd := &cobra.Command{
		Use:   `detach -i <storage_id>`,
		Short: "Detach storage volume from a server.",
		Long:  "Detach storage volume from a server.",
		Example: `  # Detach storage:
  cherryctl storage detach -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			_, err := c.Service.Detach(storageID)
			if err != nil {
				return errors.Wrap(err, "Could not detach storage from a server")
			}

			fmt.Println("Storage volume", storageID, "detached successfully.")
			return nil
		},
	}

	storageDetachCmd.Flags().IntVarP(&storageID, "storage-id", "i", 0, "The storage's ID.")

	storageDetachCmd.MarkFlagRequired("storage-id")

	return storageDetachCmd
}
