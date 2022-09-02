package storages

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	storagesListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves storage list.",
		Long:  "Retrieves a list of storages in the project.",
		Example: `  # Gets a list of storages in the specified project:
		cherryctl storage list -p 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"storage", "region", "id", "hostname"}
			storages, _, err := c.Service.List(projectID, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not get storage list")
			}

			header := []string{"ID", "Size", "Region", "Description", "Attached to"}
			data := make([][]string, len(storages))

			for i, o := range storages {
				data[i] = []string{strconv.Itoa(o.ID), fmt.Sprintf("%d %s", o.Size, o.Unit), o.Region.Name, o.Description, o.AttachedTo.Hostname}
			}

			return c.Out.Output(storages, header, &data)
		},
	}

	storagesListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	storagesListCmd.MarkFlagRequired("project-id")

	return storagesListCmd
}
