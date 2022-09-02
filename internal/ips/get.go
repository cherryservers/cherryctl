package ips

import (
	"fmt"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var ipID string
	ipGetCmd := &cobra.Command{
		Use:   `get ID`,
		Args:  cobra.ExactArgs(1),
		Short: "Get an IP address details.",
		Long:  "Get the details of the specified IP address.",
		Example: `  # Gets the details of the specified IP address:
  cherryctl ip get 30c15082-a06e-4c43-bfc3-252616b46eba`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if utils.IsValidUUID(args[0]) {
				ipID = args[0]
			} else {
				fmt.Println("IP address with ID %s was not found.", args[0])
				return nil
			}

			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"ip", "region", "hostname"}
			i, _, err := c.Service.Get(ipID, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not get IP address")
			}

			header := []string{"ID", "Address", "Cidr", "Type", "Target", "Region", "PTR record", "A record", "Tags"}
			data := make([][]string, 1)
			data[0] = []string{i.ID, i.Address, i.Cidr, i.Type, i.TargetedTo.Hostname, i.Region.Name, i.PtrRecord, i.ARecord, fmt.Sprintf("%v", i.Tags)}

			return c.Out.Output(i, header, &data)
		},
	}

	return ipGetCmd
}
