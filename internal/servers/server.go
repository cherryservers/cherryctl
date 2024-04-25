package servers

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer Servicer
	Service  cherrygo.ServersService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `server`,
		Aliases: []string{"servers", "device", "devices"},
		Short:   "Server operations. For more information on server types check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute",
		Long:    "Server operations: create, get, list, delete, start, stop, reboot, reinstall and reset-bmc-password.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Servers
		},
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
		c.Start(),
		c.Stop(),
		c.Reboot(),
		c.Reinstall(),
		c.Delete(),
		c.ResetBMC(),
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

func getServerIPByType(server cherrygo.Server, ipType string) string {
	for _, ip := range server.IPAddresses {
		if ip.Type == ipType {
			return ip.Address
		}
	}

	return ""
}
