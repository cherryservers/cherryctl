package servers

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Start() *cobra.Command {
	var serverID int
	startServerCmd := &cobra.Command{
		Use:   `start -i <server_id>`,
		Short: "Starts a server.",
		Long:  "Starts or powers on a server that is currently stopped or powered off.",
		Example: `  # Starts the specified server:
  cherryctl server start -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, _, err := c.Service.PowerOn(serverID)
			if err != nil {
				return errors.Wrap(err, "Could not start a Server")
			}

			fmt.Println("Server", serverID, "successfully started.")
			return nil
		},
	}

	startServerCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	_ = startServerCmd.MarkFlagRequired("server-id")

	return startServerCmd
}
