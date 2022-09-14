package images

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var plan string
	imageGetCmd := &cobra.Command{
		Use:     `list --plan <plan_slug>`,
		Aliases: []string{"get"},
		Short:   "Retrieves a list of images available for the given plan.",
		Long:    "Retrieves a list of images available for the given plan.",
		Example: `  # Lists the operating system images available for E5-1620v4 plan :
  cherryctl images list --plan e5_1620v4`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			images, _, err := c.Service.List(plan, c.Servicer.GetOptions())
			if err != nil {
				return errors.Wrap(err, "Could not list images")
			}

			data := make([][]string, 0)
			for _, o := range images {
				price := "0.00"
				for _, pricing := range o.Pricing {
					if pricing.Unit == "Hourly" {
						price = fmt.Sprintf("%f", pricing.Price)
						break
					}
				}

				data = append(data, []string{o.Slug, o.Name, price})
			}
			header := []string{"Slug", "Name", "Hourly price"}

			return c.Out.Output(images, header, &data)
		},
	}

	imageGetCmd.Flags().StringVarP(&plan, "plan", "", "", "The Slug or ID of a plan.")

	imageGetCmd.MarkFlagRequired("plan")

	return imageGetCmd
}
