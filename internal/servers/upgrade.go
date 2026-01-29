package servers

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) Upgrade() *cobra.Command {
	var (
		plan string
	)

	upgradeServerCmd := &cobra.Command{
		Use:   `upgrade ID --plan <plan_slug>`,
		Args:  cobra.ExactArgs(1),
		Short: "Upgrade a virtual server.",
		Long:  "Attempt to upgrade the plan of a virtual server.",
		Example: `  # Upgrade the specified server:
  cherryctl server upgrade 12345 --plan B1-2-2gb-40s-shared`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			serverID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse %q into server ID: %w", args[0], err)
			}

			_, _, err = c.Service.Upgrade(serverID, plan)
			if err != nil {
				return fmt.Errorf("failed to upgrade server %d: %w", serverID, err)
			}

			fmt.Printf("server %d upgrade has been started\n", serverID)
			return nil
		},
	}

	upgradeServerCmd.Flags().StringVar(&plan, "plan", "", "Server plan slug.")

	_ = upgradeServerCmd.MarkFlagRequired("plan")

	return upgradeServerCmd
}
