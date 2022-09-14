package storages

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer      Servicer
	Service       cherrygo.StoragesService
	ServerService cherrygo.ServersService
	Out           outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `storage`,
		Aliases: []string{"storages", "volume", "volumes"},
		Short:   "Storage operations. For more information on Elastic Block Storage check Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/",
		Long:    "Storage operations: create, get, list, delete, attach and detach.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Storages
			c.ServerService = c.Servicer.API(cmd).Servers
		},
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
		c.Attach(),
		c.Detach(),
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
