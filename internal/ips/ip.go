package ips

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer      Servicer
	Service       cherrygo.IpAddressesService
	ServerService cherrygo.ServersService
	Out           outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `ip`,
		Aliases: []string{"ips", "ip-address", "ip-addresses"},
		Short:   "IP address operations. For more information on IP addresses, check out the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing",
		Long:    "IP address operations: get, list, create, update, assign, unassign and delete.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).IPAddresses
			c.ServerService = c.Servicer.API(cmd).Servers
		},
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
		c.Assign(),
		c.Unassign(),
		c.Delete(),
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
