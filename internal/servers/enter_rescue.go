package servers

import (
	"fmt"
	"github.com/cherryservers/cherrygo/v3"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) EnterRescue() *cobra.Command {
	var password string

	enterRescueServerCmd := &cobra.Command{
		Use:   `enter-rescue ID --password <password>`,
		Args:  cobra.ExactArgs(1),
		Short: "Enter server rescue mode.",
		Long:  "Put the specified server in rescue mode.",
		Example: `  # Put the specified server in rescue mode:
  cherryctl server enter-rescue 12345 --password abcdef`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			request := &cherrygo.RescueServerFields{
				Password: password,
			}

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.EnterRescueMode(serverID, request)
				if err != nil {
					return errors.Wrap(err, "Couldn't put server in rescue mode.")
				}

				fmt.Println("Server", serverID, "successfully entered rescue mode.")
				return nil
			}

			fmt.Println("Server with ID %s was not found", args[0])
			return nil
		},
	}

	enterRescueServerCmd.Flags().StringVarP(&password, "password", "", "", "Server password.")

	_ = enterRescueServerCmd.MarkFlagRequired("password")

	return enterRescueServerCmd
}
