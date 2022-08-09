package projects

import (
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		projectID int
		bgp       bool
		name      string
	)
	projectUpdateCmd := &cobra.Command{
		Use:   `update [-p <project_id>] [--name <project_name>] [--bgp <bool>]`,
		Short: "Update a project.",
		Long:  "Update a project.",
		Example: `  # Update project to enable BGP:
  cherryctl project update -p 12345 --bgp true`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.UpdateProject{
				Bgp: &bgp,
			}

			if name != "" {
				request.Name = &name
			}

			o, _, err := c.Service.Update(projectID, request)
			if err != nil {
				return errors.Wrap(err, "Could not update project")
			}

			header := []string{"ID", "Name", "BGP enabled", "BGP ASN"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Name, utils.BoolToYesNo(o.Bgp.Enabled), strconv.Itoa(o.Bgp.LocalASN)}

			return c.Out.Output(o, header, &data)
		},
	}

	projectUpdateCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	projectUpdateCmd.Flags().BoolVarP(&bgp, "bgp", "b", false, "True to enable BGP in a project.")
	projectUpdateCmd.Flags().StringVarP(&name, "name", "", "", "Project name.")

	projectUpdateCmd.MarkFlagRequired("project-id")

	return projectUpdateCmd
}
