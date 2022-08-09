package storages

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var storageID int
	storagesGetCmd := &cobra.Command{
		Use:   `get [-i <storage_id>]`,
		Short: "Retrieves storage details.",
		Long:  "Retrieves the details of the specified storage.",
		Example: `  # Gets the details of the specified storage:
  cherryctl storage get -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"storage", "region", "id", "hostname"}
			o, _, err := c.Service.Get(storageID, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not get storage")
			}
			header := []string{"ID", "Size", "Region", "Description", "Attached to"}

			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), fmt.Sprintf("%d %s", o.Size, o.Unit), o.Region.Name, o.Description, o.AttachedTo.Hostname}

			return c.Out.Output(o, header, &data)
		},
	}

	storagesGetCmd.Flags().IntVarP(&storageID, "storage-id", "i", 0, "The ID of a storage volume.")
	_ = storagesGetCmd.MarkFlagRequired("storage-id")

	return storagesGetCmd
}
