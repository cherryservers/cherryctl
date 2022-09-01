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
		Use:   `get ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Retrieves storage details.",
		Long:  "Retrieves the details of the specified storage.",
		Example: `  # Gets the details of the specified storage:
  cherryctl storage get 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if storID, err := strconv.Atoi(args[0]); err == nil {
				storageID = storID
			}
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

	return storagesGetCmd
}
