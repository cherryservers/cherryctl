package teams

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		teamID   int
		currency string
		name     string
		teamType string
	)
	teamUpdateCmd := &cobra.Command{
		Use:   `update [-t <team_id>] [--name <team_name>] [--currency <currency_code>] [--type <team_type>]`,
		Short: "Update a team.",
		Long:  "Update a team.",
		Example: `  # Update a team to change currency to EUR:
  cherryctl team update -t 12345 --currency EUR`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.UpdateTeam{}

			if currency != "" {
				request.Currency = &currency
			}

			if name != "" {
				request.Name = &name
			}

			if teamType != "" {
				request.Type = &teamType
			}

			o, _, err := c.Service.Update(teamID, request)
			if err != nil {
				return errors.Wrap(err, "Could not update team")
			}

			header := []string{"ID", "Name", "Remaining credit", "Hourly usage", "Currency"}
			data := make([][]string, 1)
			credit := o.Credit.Account.Remaining + o.Credit.Promo.Remaining
			data[0] = []string{strconv.Itoa(o.ID), o.Name, fmt.Sprintf("%f", credit), fmt.Sprintf("%f", o.Credit.Resources.Pricing.Price), o.Billing.Currency}

			return c.Out.Output(o, header, &data)
		},
	}

	teamUpdateCmd.Flags().IntVarP(&teamID, "team-id", "t", 0, "The team's ID.")
	teamUpdateCmd.Flags().StringVarP(&currency, "currency", "", "", "Team currency, available otions: EUR, USD.")
	teamUpdateCmd.Flags().StringVarP(&teamType, "type", "", "", "Team type, available options: personal, business.")
	teamUpdateCmd.Flags().StringVarP(&name, "name", "", "", "Team name.")

	teamUpdateCmd.MarkFlagRequired("team-id")

	return teamUpdateCmd
}
