package storages

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		storageID   int
		size        int
		description string
	)
	storageUpdateCmd := &cobra.Command{
		Use:   `update ID [--size <gigabytes>] [--description <text>]`,
		Args:  cobra.ExactArgs(1),
		Short: "Update storage volume.",
		Long:  "Update storage size or description.",
		Example: `  # Update storage size to 1000 gigabyte:
  cherryctl storage update 12345 --size 1000`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if storID, err := strconv.Atoi(args[0]); err == nil {
				storageID = storID
			}
			request := &cherrygo.UpdateStorage{
				StorageID:   storageID,
				Size:        size,
				Description: description,
			}

			o, _, err := c.Service.Update(request)
			if err != nil {
				return errors.Wrap(err, "Could not update storage")
			}

			header := []string{"ID", "Size", "Region", "Description", "Attached to"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), fmt.Sprintf("%d %s", o.Size, o.Unit), o.Region.Name, o.Description, o.AttachedTo.Hostname}

			return c.Out.Output(o, header, &data)
		},
	}

	storageUpdateCmd.Flags().IntVarP(&size, "size", "", 0, "Storage volume size in gigabytes. Value must be greater than current volume size.")
	storageUpdateCmd.Flags().StringVarP(&description, "description", "", "", "Storage description.")

	return storageUpdateCmd
}
