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
	var hostname string
	var projectID int
	serverGetCmd := &cobra.Command{
		Use:   `get {-i <server_id> | --hostname} [-p <project_id>]`,
		Short: "Retrieves server details.",
		Long:  "Retrieves the details of the specified server.",
		Example: `  # Gets the details of the specified server:
  cherryctl server get -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if hostname == "" && serverID == 0 {
				return fmt.Errorf("either server-id or hostname should be set")
			}
			if hostname != "" {
				srvID, err := utils.ServerHostnameToID(hostname, projectID, c.Service)
				if err != nil {
					return errors.Wrap(err, "Could not get a Server")
				}
				serverID = srvID
			}

			s, _, err := c.Service.Get(serverID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get Server")
			}
			header := []string{"ID", "Plan", "Hostname", "Image", "State", "Public IP", "Private IP", "Region", "Tags", "Spot"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(s.ID), s.Plan.Name, s.Hostname, s.Image, s.State, getServerIPByType(s, "primary-ip"), getServerIPByType(s, "private-ip"), s.Region.Name, fmt.Sprintf("%v", s.Tags), utils.BoolToYesNo(s.SpotInstance)}

			return c.Out.Output(s, header, &data)
		},
	}

	serverGetCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	serverGetCmd.Flags().StringVarP(&hostname, "hostname", "", "", "The Hostname of a server.")
	serverGetCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")

	serverGetCmd.MarkFlagsMutuallyExclusive("server-id", "hostname")

	return serverGetCmd
}
