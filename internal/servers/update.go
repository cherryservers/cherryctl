package servers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		serverID int
		tags     []string
		name     string
		hostname string
		bgp      bool
	)
	serverUpdateCmd := &cobra.Command{
		Use:   `update ID [--name <server_name>] [--hostname <hostname>] [--tags <tags>] [--bgp]`,
		Args:  cobra.ExactArgs(1),
		Short: "Update server.",
		Long:  "Update server.",
		Example: `  # Update server to change tags:
  cherryctl server update 12345 --tags="env=stage"`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if srvID, err := strconv.Atoi(args[0]); err == nil {
				serverID = srvID
			} else {
				fmt.Println("Server with ID %s was not found.", args[0])
				return nil
			}

			tagsArr := make(map[string]string)

			for _, kv := range tags {
				var k, v string
				tokens := strings.SplitN(kv, "=", 2)
				k = strings.TrimSpace(tokens[0])
				if len(tokens) != 1 {
					v = strings.TrimSpace(tokens[1])
				}

				tagsArr[k] = v
			}

			request := &cherrygo.UpdateServer{
				Name:     name,
				Hostname: hostname,
				Bgp:      bgp,
				Tags:     &tagsArr,
			}

			s, _, err := c.Service.Update(serverID, request)
			if err != nil {
				return errors.Wrap(err, "Could not update server")
			}

			header := []string{"ID", "Name", "Hostname", "Image", "State", "Region", "Tags"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(s.ID), s.Name, s.Hostname, s.Image, s.State, s.Region.Name, utils.FormatStringTags(&s.Tags)}

			return c.Out.Output(s, header, &data)
		},
	}

	serverUpdateCmd.Flags().StringVarP(&hostname, "hostname", "", "", "Server hostname.")
	serverUpdateCmd.Flags().BoolVarP(&bgp, "bgp", "b", false, "True to enable BGP in a server.")
	serverUpdateCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, `Tag or list of tags for the server: --tags="key=value,env=prod".`)
	serverUpdateCmd.Flags().StringVarP(&name, "name", "", "", "Server name.")

	return serverUpdateCmd
}
