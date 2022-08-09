package servers

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Stop() *cobra.Command {
	var serverID int
	stopServerCmd := &cobra.Command{
		Use:   `stop -i <server_id>`,
		Short: "Stop a server.",
		Long:  "Stops or powers off a server that is currently powered on.",
		Example: `  # Stops the specified server:
  cherryctl server stop -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, _, err := c.Service.PowerOff(serverID)
			if err != nil {
				return errors.Wrap(err, "Could not stop a Server")
			}

			fmt.Println("Server", serverID, "successfully stopped.")
			return nil
		},
	}

	stopServerCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	_ = stopServerCmd.MarkFlagRequired("server-id")

	return stopServerCmd
}
