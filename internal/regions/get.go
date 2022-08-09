package regions

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var regionID string
	regionGetCmd := &cobra.Command{
		Use:   `get [-i <region_slug>]`,
		Short: "Retrieves region details.",
		Long:  "Retrieves the details of the specified region.",
		Example: `  # Gets the details of the specified region:
  cherryctl region get -i eu_nord_1`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			o, _, err := c.Service.Get(regionID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get region")
			}

			header := []string{"ID", "Slug", "Name", "BGP hosts", "BGP asn"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Slug, o.Name, strings.Join(o.BGP.Hosts, ", "), strconv.Itoa(o.BGP.Asn)}

			return c.Out.Output(o, header, &data)
		},
	}

	regionGetCmd.Flags().StringVarP(&regionID, "region-id", "i", "", "The Slug or ID of region.")
	_ = regionGetCmd.MarkFlagRequired("region-id")

	return regionGetCmd
}
