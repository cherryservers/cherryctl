package projects

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var projectID int
	projectGetCmd := &cobra.Command{
		Use:   `get ID [-p <project_id>]`,
		Short: "Retrieves project details.",
		Long:  "Retrieves the details of the specified project.",
		Example: `  # Gets the details of the specified project:
  cherryctl project get 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if len(args) > 0 {
				prID, err := strconv.Atoi(args[0])
				if err == nil {
					projectID = prID
				}
			}

			if projectID == 0 {
				return fmt.Errorf("project-id should be set %v\t", projectID)
			}

			o, _, err := c.Service.Get(projectID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get project")
			}

			header := []string{"ID", "Name", "BGP enabled", "BGP ASN"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Name, utils.BoolToYesNo(o.Bgp.Enabled), strconv.Itoa(o.Bgp.LocalASN)}

			return c.Out.Output(o, header, &data)
		},
	}

	projectGetCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")

	return projectGetCmd
}
