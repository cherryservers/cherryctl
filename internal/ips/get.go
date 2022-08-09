package ips

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var ipID string
	ipGetCmd := &cobra.Command{
		Use:   `get [-i <ip_address_id>]`,
		Short: "Get an IP address details.",
		Long:  "Get the details of the specified IP address.",
		Example: `  # Gets the details of the specified IP address:
  cherryctl ip get -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
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

	ipGetCmd.Flags().StringVarP(&ipID, "ip-address-id", "i", "", "The ID of a IP address.")
	_ = ipGetCmd.MarkFlagRequired("ip-address-id")

	return ipGetCmd
}
