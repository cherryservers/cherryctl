package regions

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer Servicer
	Service  cherrygo.RegionsService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `region`,
		Aliases: []string{"regions", "location", "locations"},
		Short:   "Region operations. For more information on Networking, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#network",
		Long:    "Region operations: get, list.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Regions
		},
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
	)

	return cmd
}

type Servicer interface {
	API(*cobra.Command) *cherrygo.Client
	GetOptions() *cherrygo.GetOptions
	Config(cmd *cobra.Command) *viper.Viper
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
