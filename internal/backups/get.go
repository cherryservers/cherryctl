package backups

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var storageID int
	backupGetCmd := &cobra.Command{
		Use:   `get <backup_ID>`,
		Args:  cobra.ExactArgs(1),
		Short: "Retrieves backup storage details.",
		Long:  "Retrieves the details of the specified backup storage.",
		Example: `  # Gets the details of the specified backup storage:
  cherryctl backup get 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if storID, err := strconv.Atoi(args[0]); err == nil {
				storageID = storID
			}

			o, _, err := c.Service.Get(storageID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get a backup storage")
			}
			header := []string{"ID", "Status", "Attached to", "Size GB", "Used GB", "Region", "Price"}

			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Status, o.AttachedTo.Hostname, strconv.Itoa(o.SizeGigabytes), strconv.Itoa(o.UsedGigabytes), o.Region.Slug, fmt.Sprint(o.Pricing.Price)}

			return c.Out.Output(o, header, &data)
		},
	}

	return backupGetCmd
}
