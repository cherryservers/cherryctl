package servers

import (
	"fmt"
	"github.com/cherryservers/cherrygo/v3"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Rescue() *cobra.Command {
	var password string

	rescueServerCmd := &cobra.Command{
		Use:   `rescue ID --password <password>`,
		Args:  cobra.ExactArgs(1),
		Short: "Rescue a server.",
		Long:  "Rescue the specified server.",
		Example: `  # Rescue the specified server:
  cherryctl server rescue 12345 --password abcdef`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			request := &cherrygo.RescueServerFields{
				Password: password,
			}

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.Rescue(serverID, request)
				if err != nil {
					return errors.Wrap(err, "Could not rescue a Server")
				}

				fmt.Println("Server", serverID, "successfully entered rescue mode.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}

	rescueServerCmd.Flags().StringVarP(&password, "password", "", "", "Server password.")
	_ = rescueServerCmd.MarkFlagRequired("password")

	return rescueServerCmd
}
