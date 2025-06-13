package projects

import (
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		teamID int
		bgp    bool
		name   string
	)
	projectCreateCmd := &cobra.Command{
		Use:   `create -t <team_id> --name <project_name> [--bgp <bool>]`,
		Short: "Create a project.",
		Long:  "Create a new project for the specified team.",
		Example: `  # To create a new project with BGP support enabled:
  cherryctl project create -t 12345 --name "Project with BGP" --bgp`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.CreateProject{
				Name: name,
				Bgp:  bgp,
			}

			o, _, err := c.Service.Create(teamID, request)
			if err != nil {
				return errors.Wrap(err, "Could not create project")
			}

			header := []string{"ID", "Name", "BGP enabled", "BGP ASN"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Name, utils.BoolToYesNo(o.Bgp.Enabled), strconv.Itoa(o.Bgp.LocalASN)}

			return c.Out.Output(o, header, &data)
		},
	}

	projectCreateCmd.Flags().IntVarP(&teamID, "team-id", "t", 0, "The teams's ID.")
	projectCreateCmd.Flags().BoolVarP(&bgp, "bgp", "b", false, "Enable BGP support.")
	projectCreateCmd.Flags().StringVarP(&name, "name", "", "", "Project name.")

	projectCreateCmd.MarkFlagRequired("team-id")
	projectCreateCmd.MarkFlagRequired("name")

	return projectCreateCmd
}
