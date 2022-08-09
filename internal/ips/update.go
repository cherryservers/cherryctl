package ips

import (
	"fmt"
	"strings"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		ipID      string
		ptrRecord string
		aRecord   string
		tags      []string
	)
	ipUpdateCmd := &cobra.Command{
		Use:   `update -i <ip_address_id> [--ptr-record] [--a-record] [--tags]`,
		Short: "Update IP address.",
		Long:  "Update tags, ptr record, a record or target server of a IP address.",
		Example: `  # Updates a record and tags:
  cherryctl ip update -i 30c15082-a06e-4c43-bfc3-252616b46eba --a-record stage --tags="env=stage"`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			tagsArr := make(map[string]string)
			for _, kv := range tags {
				var k, v string
				tokens := strings.SplitN(kv, "=", 2)
				k = strings.TrimSpace(tokens[0])
				if len(tokens) != 1 {
					v = strings.TrimSpace(tokens[1])
				}

				tagsArr[k] = v
			}

			request := &cherrygo.UpdateIPAddress{
				PtrRecord: ptrRecord,
				ARecord:   aRecord,
			}

			if len(tagsArr) > 0 {
				request.Tags = &tagsArr
			}

			i, _, err := c.Service.Update(ipID, request)
			if err != nil {
				return errors.Wrap(err, "Could not update IP address")
			}

			header := []string{"ID", "Address", "Cidr", "Type", "Region", "PTR record", "A record", "Tags"}
			data := make([][]string, 1)
			data[0] = []string{i.ID, i.Address, i.Cidr, i.Type, i.Region.Name, i.PtrRecord, i.ARecord, fmt.Sprintf("%v", i.Tags)}

			return c.Out.Output(i, header, &data)
		},
	}

	ipUpdateCmd.Flags().StringVarP(&ipID, "ip-address-id", "i", "", "The ID of a IP address.")
	ipUpdateCmd.Flags().StringVarP(&ptrRecord, "ptr-record", "", "", "Slug of the region from where IP address will requested.")
	ipUpdateCmd.Flags().StringVarP(&aRecord, "a-record", "", "", "Slug of the region from where IP address will requested.")
	ipUpdateCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, `Tag or list of tags for the server: --tags="key=value,env=prod".`)

	ipUpdateCmd.MarkFlagRequired("ip-address-id")

	return ipUpdateCmd
}
