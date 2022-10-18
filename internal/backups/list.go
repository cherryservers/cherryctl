package backups

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	backupsListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves a list of backup storages.",
		Long:  "Retrieves a list of backup storages in the project.",
		Example: `  # Gets a list of backup storages in the specified project:
		cherryctl backup list -p 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			storages, _, err := c.Service.ListBackups(projectID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get backup storage list")
			}

			header := []string{"ID", "Status", "Attached to", "Size GB", "Used GB", "Region", "Price"}
			data := make([][]string, len(storages))

			for i, o := range storages {
				data[i] = []string{strconv.Itoa(o.ID), o.Status, o.AttachedTo.Hostname, strconv.Itoa(o.SizeGigabytes), strconv.Itoa(o.UsedGigabytes), o.Region.Slug, fmt.Sprint(o.Pricing.Price)}
			}

			return c.Out.Output(storages, header, &data)
		},
	}

	backupsListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	backupsListCmd.MarkFlagRequired("project-id")

	return backupsListCmd
}
