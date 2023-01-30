package ips

import (
	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	var types []string
	ipListCmd := &cobra.Command{
		Use:   `list -p <project_id>`,
		Short: "Retrieves list of IP addresses.",
		Long:  "Retrieves the details of IP addresses in the project.",
		Example: `  # Gets a list of IP addresses in the specified project:
  cherryctl ip list -p 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"ip", "region", "hostname"}

			if len(types) > 0 {
				getOptions.Type = types
			}

			ips, _, err := c.Service.List(projectID, getOptions)
			if err != nil {
				return errors.Wrap(err, "Could not list servers")
			}
			data := make([][]string, len(ips))

			for i, ip := range ips {
				data[i] = []string{ip.ID, ip.Address, ip.Cidr, ip.Type, ip.TargetedTo.Hostname, ip.Region.Name, ip.Region.Name, ip.PtrRecord, utils.FormatStringTags(ip.Tags)}
			}
			header := []string{"ID", "Address", "Cidr", "Type", "Target", "Region", "PTR record", "A record", "Tags"}

			return c.Out.Output(ips, header, &data)
		},
	}

	ipListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	ipListCmd.Flags().StringSliceVarP(&types, "type", "", []string{}, `Comma separated list of available IP addresses types (subnet,primary-ip,floating-ip,private-ip)`)

	ipListCmd.MarkFlagRequired("project-id")

	return ipListCmd
}
