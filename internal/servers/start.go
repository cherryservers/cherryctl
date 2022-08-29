package servers

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Start() *cobra.Command {
	startServerCmd := &cobra.Command{
		Use:   `start ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Starts a server.",
		Long:  "Starts or powers on a server that is currently stopped or powered off.",
		Example: `  # Starts the specified server:
  cherryctl server start 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.PowerOn(serverID)
				if err != nil {
					return errors.Wrap(err, "Could not start a Server")
				}

				fmt.Println("Server", serverID, "successfully started.")
				return nil
			}

			fmt.Println("Server with ID %s was not found.", args[0])
			return nil
		},
	}

	return startServerCmd
}
