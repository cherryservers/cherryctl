package storages

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Detach() *cobra.Command {
	var (
		storageID int
	)
	storageDetachCmd := &cobra.Command{
		Use:   `detach ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Detach storage volume from a server.",
		Long:  "Detach storage volume from a server.",
		Example: `  # Detach storage:
  cherryctl storage detach 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if storID, err := strconv.Atoi(args[0]); err == nil {
				storageID = storID
			}

			_, err := c.Service.Detach(storageID)
			if err != nil {
				return errors.Wrap(err, "Could not detach storage from a server")
			}

			fmt.Println("Storage volume", storageID, "detached successfully.")
			return nil
		},
	}

	return storageDetachCmd
}
