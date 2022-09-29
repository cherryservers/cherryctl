package teams

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var teamID int
	teamGetCmd := &cobra.Command{
		Use:   `get ID`,
		Short: "Retrieves team details.",
		Long:  "Retrieves the details of the specified team.",
		Example: `  # Gets the details of the specified team:
  cherryctl team get 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if len(args) > 0 {
				tID, err := strconv.Atoi(args[0])
				if err == nil {
					teamID = tID
				}
			}

			if teamID == 0 {
				return fmt.Errorf("team-id should be set %v\t", teamID)
			}

			o, _, err := c.Service.Get(teamID, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not get team")
			}

			header := []string{"ID", "Name", "Remaining credit", "Hourly usage", "Currency"}
			data := make([][]string, 1)
			credit := o.Credit.Account.Remaining + o.Credit.Promo.Remaining
			data[0] = []string{strconv.Itoa(o.ID), o.Name, fmt.Sprintf("%f", credit), fmt.Sprintf("%f", o.Credit.Resources.Pricing.Price), o.Billing.Currency}

			return c.Out.Output(o, header, &data)
		},
	}

	teamGetCmd.Flags().IntVarP(&teamID, "team-id", "i", 0, "The team's ID.")

	return teamGetCmd
}
