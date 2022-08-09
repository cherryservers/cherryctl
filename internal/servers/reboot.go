package servers

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Reboot() *cobra.Command {
	var serverID int
	rebootServerCmd := &cobra.Command{
		Use:   `reboot -i <server_id>`,
		Short: "Reboot a server.",
		Long:  "Reboot the specified server.",
		Example: `  # Reboot the specified server:
  cherryctl server reboot -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, _, err := c.Service.Reboot(serverID)
			if err != nil {
				return errors.Wrap(err, "Could not reboot a Server")
			}

			fmt.Println("Server", serverID, "successfully rebooted.")
			return nil
		},
	}

	rebootServerCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	_ = rebootServerCmd.MarkFlagRequired("server-id")

	return rebootServerCmd
}
