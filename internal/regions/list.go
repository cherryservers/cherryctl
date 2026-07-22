package regions

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	regionListCmd := &cobra.Command{
		Use:   `list`,
		Short: "Retrieves list of regions.",
		Long:  "Retrieves list of regions.",
		Example: `  # Gets list of regions:
  cherryctl region list`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			ctx := cmd.Context()

			list, _, err := c.Service.List(ctx, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get regions list")
			}

			data := make([][]string, len(list))
			header := []string{"ID", "Slug", "Name", "BGP hosts", "BGP asn"}

			for i, o := range list {
				data[i] = []string{strconv.Itoa(o.ID), o.Slug, o.Name, strings.Join(o.BGP.Hosts, ", "), strconv.Itoa(o.BGP.ASN)}
			}

			return c.Out.Output(list, header, &data)
		},
	}

	return regionListCmd
}
