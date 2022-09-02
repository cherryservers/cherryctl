package servers

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Reboot() *cobra.Command {
	rebootServerCmd := &cobra.Command{
		Use:   `reboot ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Reboot a server.",
		Long:  "Reboot the specified server.",
		Example: `  # Reboot the specified server:
  cherryctl server reboot 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.Reboot(serverID)
				if err != nil {
					return errors.Wrap(err, "Could not reboot a Server")
				}

				fmt.Println("Server", serverID, "successfully rebooted.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}

	return rebootServerCmd
}
