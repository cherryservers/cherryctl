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
		Use:   `get {ID | SLUG}`,
		Args:  cobra.ExactArgs(1),
		Short: "Retrieves region details.",
		Long:  "Retrieves the details of the specified region.",
		Example: `  # Gets the details of the specified region:
  cherryctl region get LT-Siauliai`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			regionID = args[0]
			o, _, err := c.Service.Get(regionID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get a region")
			}

			header := []string{"ID", "Slug", "Name", "BGP hosts", "BGP asn"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Slug, o.Name, strings.Join(o.BGP.Hosts, ", "), strconv.Itoa(o.BGP.Asn)}

			return c.Out.Output(o, header, &data)
		},
	}

	return regionGetCmd
}
