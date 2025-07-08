package servers

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) ListCycles() *cobra.Command {
	serverListCmd := &cobra.Command{
		Use:   `list-cycles`,
		Short: "Retrieves server billing cycle list.",
		Long:  "Retrieves a list of possible server billing cycles.",
		Example: `  # Gets a list of server billing cycles:
		cherryctl server list-cycles`,

		RunE: func(cmd *cobra.Command, args []string) error {
			options := c.Servicer.GetOptions()

			cmd.SilenceUsage = true
			cycles, _, err := c.Service.ListCycles(options)
			if err != nil {
				return errors.Wrap(err, "Could not list server billing cycles.")
			}
			data := make([][]string, len(cycles))

			for i, s := range cycles {
				data[i] = []string{strconv.Itoa(s.ID), s.Name, s.Slug}
			}
			header := []string{"ID", "Name", "Slug"}

			return c.Out.Output(cycles, header, &data)
		},
	}

	return serverListCmd
}
