package servers

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	var search string
	serverListCmd := &cobra.Command{
		Use:     `list -p <project_id>`,
		Aliases: []string{"list"},
		Short:   "Retrieves server list.",
		Long:    "Retrieves a list of servers in the project.",
		Example: `  # Gets a list of servers in the specified project:
		cherryctl server list -p 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if projectID == 0 {
				return fmt.Errorf("project-id should be set %v\t", projectID)
			}

			options := c.Servicer.GetOptions()

			if search != "" {
				options.QueryParams = map[string]string{"search": search}
			}

			cmd.SilenceUsage = true
			servers, _, err := c.Service.List(projectID, options)
			if err != nil {
				return errors.Wrap(err, "Could not list servers")
			}
			data := make([][]string, len(servers))

			for i, s := range servers {
				data[i] = []string{strconv.Itoa(s.ID), s.Name, s.Hostname, s.Image, s.State, getServerIPByType(s, "primary-ip"), getServerIPByType(s, "private-ip"), s.Region.Name, utils.FormatStringTags(&s.Tags), utils.BoolToYesNo(s.SpotInstance)}
			}
			header := []string{"ID", "Name", "Hostname", "Image", "State", "Public IP", "Private IP", "Region", "Tags", "Spot"}

			return c.Out.Output(servers, header, &data)
		},
	}

	serverListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	serverListCmd.Flags().StringVarP(&search, "search", "", "", "Search server by Hostname or Public IP phrase.")

	serverListCmd.MarkFlagRequired("project-id")

	return serverListCmd
}
