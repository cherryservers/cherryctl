package projects

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var teamID int
	projectListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves a list of projects details.",
		Long:  "Retrieves the details of projects.",
		Example: `  # List available projects:
  cherryctl project list -t 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if teamID == 0 {
				return fmt.Errorf("team-id should be set %v\t", teamID)
			}

			projects, _, err := c.Service.List(teamID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get projects list")
			}

			data := make([][]string, len(projects))
			for i, o := range projects {
				data[i] = []string{strconv.Itoa(o.ID), o.Name, utils.BoolToYesNo(o.Bgp.Enabled), strconv.Itoa(o.Bgp.LocalASN)}
			}
			header := []string{"ID", "Name", "BGP enabled", "BGP ASN"}

			return c.Out.Output(projects, header, &data)
		},
	}

	projectListCmd.Flags().IntVarP(&teamID, "team-id", "t", 0, "The team's ID.")
	_ = projectListCmd.MarkFlagRequired("team-id")

	return projectListCmd
}
