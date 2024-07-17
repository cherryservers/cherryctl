package servers

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) ExitRescue() *cobra.Command {
	exitRescueServerCmd := &cobra.Command{
		Use:   `exit-rescue ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Exit server rescue mode.",
		Long:  "Put the specified server out of rescue mode.",
		Example: `  # Put the specified server out of rescue mode:
  cherryctl server exit-rescue 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.ExitRescueMode(serverID)
				if err != nil {
					return errors.Wrap(err, "Could not put server out of rescue mode.")
				}

				fmt.Println("Server", serverID, "successfully exited rescue mode.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}
	return exitRescueServerCmd
}
