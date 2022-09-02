package servers

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var serverID int
	var projectID int
	serverGetCmd := &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Server ID or Hostname must be given as positional argument")
			}
			return nil
		},
		Use:   `get {ID | HOSTNAME} [-p <project_id>]`,
		Short: "Retrieves server details.",
		Long:  "Retrieves the details of the specified server.",
		Example: `  # Gets the details of the specified server:
  cherryctl server get 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if srvID, err := strconv.Atoi(args[0]); err == nil {
				serverID = srvID
			} else {
				srvID, err := utils.ServerHostnameToID(args[0], projectID, c.Service)
				if err != nil {
					return errors.Wrap(err, "Server with hostname %s was not found")
				}
				serverID = srvID
			}

			s, _, err := c.Service.Get(serverID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get a Server")
			}
			header := []string{"ID", "Plan", "Hostname", "Image", "State", "Public IP", "Private IP", "Region", "Tags", "Spot"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(s.ID), s.Plan.Name, s.Hostname, s.Image, s.State, getServerIPByType(s, "primary-ip"), getServerIPByType(s, "private-ip"), s.Region.Name, fmt.Sprintf("%v", s.Tags), utils.BoolToYesNo(s.SpotInstance)}

			return c.Out.Output(s, header, &data)
		},
	}

	serverGetCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")

	return serverGetCmd
}
