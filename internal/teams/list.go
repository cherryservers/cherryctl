package teams

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	teamListCmd := &cobra.Command{
		Use:     `list`,
		Aliases: []string{"ls"},
		Short:   "Retrieves list of teams details.",
		Long:    "Retrieves the details of teams.",
		Example: `  # List available teams:
  cherryctl team list`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			teams, _, err := c.Service.List(c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get teams list")
			}

			data := make([][]string, len(teams))
			for i, o := range teams {
				credit := o.Credit.Account.Remaining + o.Credit.Promo.Remaining
				data[i] = []string{strconv.Itoa(o.ID), o.Name, fmt.Sprintf("%f", credit), fmt.Sprintf("%f", o.Credit.Resources.Pricing.Price), o.Billing.Currency}
			}
			header := []string{"ID", "Name", "Remaining credit", "Hourly usage", "Currency"}

			return c.Out.Output(teams, header, &data)
		},
	}

	return teamListCmd
}
