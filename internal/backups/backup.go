package backups

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer      Servicer
	Service       cherrygo.BackupsService
	ServerService cherrygo.ServersService
	Out           outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `backup`,
		Aliases: []string{"backups"},
		Short:   "Server backup operations. For more information on backups check Product Docs: https://docs.cherryservers.com/knowledge/backup-storage",
		Long:    "Server backup storage operations: plans, get, list, create, update, methods, update-method, method-whitelist and remove.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Backups
			c.ServerService = c.Servicer.API(cmd).Servers
		},
	}

	cmd.AddCommand(
		c.Plans(),
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
		c.Upgrade(),
		c.Methods(),
		c.MethodsWhitelist(),
		c.UpdateMethod(),
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
