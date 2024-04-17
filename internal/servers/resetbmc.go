package servers

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
)

func (c *Client) ResetBMC() *cobra.Command {
	resetServerBMCCmd := &cobra.Command{
		Use:   `reset-bmc-password ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Reset server BMC password.",
		Long:  "Reset BMC password for the specified server.",
		Example: `  # Reset BMC password for the specified server:
  cherryctl server reset-bmc-password 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.ResetBMCPassword(serverID)
				if err != nil {
					return errors.Wrap(err, "Could not reset server BMC password")
				}

				fmt.Println("Server", serverID, "BMC password successfully reset.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}

	return resetServerBMCCmd
}
