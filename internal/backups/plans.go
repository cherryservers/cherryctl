package backups

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Plans() *cobra.Command {
	backupPlansCmd := &cobra.Command{
		Use:     `plans`,
		Aliases: []string{"plan"},
		Short:   "Retrieves available backup storage plans.",
		Long:    "Retrieves the details of available backup storage plans.",
		Example: `  # Gets the list of available backup storage plans:
  cherryctl backup plans`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			plans, _, err := c.Service.ListPlans(c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get backup storage plan list")
			}
			header := []string{"Name", "Slug", "Size GB", "Region", "Price"}
			data := make([][]string, 0)
			for _, p := range plans {
				priceHour := "-"
				for _, pricing := range p.Pricing {
					if pricing.Unit == "Hourly" {
						priceHour = fmt.Sprintf("%f", pricing.Price)
					}
				}
				for _, r := range p.Regions {
					data = append(data, []string{p.Name, p.Slug, strconv.Itoa(p.SizeGigabytes), r.Slug, priceHour})
				}
			}

			return c.Out.Output(plans, header, &data)
		},
	}

	return backupPlansCmd
}
