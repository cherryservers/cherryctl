package teams

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		currency       string = "EUR"
		name, teamType string
	)
	teamCreateCmd := &cobra.Command{
		Use:   `create --name <team_name> --type <team_type> [--currency <currency_code>]`,
		Short: "Create a team.",
		Long:  "Create a team.",
		Example: `  # Create business team with USD currency:
  cherryctl team create --name="USD team" --type business --currency USD`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.CreateTeam{
				Name:     name,
				Type:     teamType,
				Currency: currency,
			}

			o, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not create a Team")
			}

			header := []string{"ID", "Name", "Remaining credit", "Hourly usage", "Currency"}
			data := make([][]string, 1)
			credit := o.Credit.Account.Remaining + o.Credit.Promo.Remaining
			data[0] = []string{strconv.Itoa(o.ID), o.Name, fmt.Sprintf("%f", credit), fmt.Sprintf("%f", o.Credit.Resources.Pricing.Price), o.Billing.Currency}

			return c.Out.Output(o, header, &data)
		},
	}

	teamCreateCmd.Flags().StringVarP(&name, "name", "", "", "Team name.")
	teamCreateCmd.Flags().StringVarP(&teamType, "type", "", "", "Team type, available options: personal, business.")
	teamCreateCmd.Flags().StringVarP(&currency, "currency", "", "", "Team currency, available otions: EUR, USD.")

	teamCreateCmd.MarkFlagRequired("type")
	teamCreateCmd.MarkFlagRequired("name")

	return teamCreateCmd
}
