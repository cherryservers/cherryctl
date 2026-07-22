package servers

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
)

type Client struct {
	Deps
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `server`,
		Aliases: []string{"servers", "device", "devices"},
		Short:   "Server operations. For more information on server types, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute",
		Long:    "Server operations: create, get, list, delete, start, stop, reboot, reinstall, reset-bmc-password and list-cycles.",
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
		c.Start(),
		c.Stop(),
		c.Reboot(),
		c.EnterRescue(),
		c.ExitRescue(),
		c.Reinstall(),
		c.Delete(),
		c.ResetBMC(),
		c.ListCycles(),
		c.Upgrade(),
	)

	return cmd
}

type Deps interface {
	Client() cherrygo.ServersService
	GetOpts() *cherrygo.GetOptions
	Outputer() outputs.Outputer
}

func NewClient(dep Deps) *Client {
	return &Client{
		Deps: dep,
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
