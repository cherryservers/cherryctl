package users

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer Servicer
	Service  cherrygo.UsersService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `user`,
		Aliases: []string{"users"},
		Short:   "User operations.",
		Long:    "User operations: get.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).Users
		},
	}

	cmd.AddCommand(
		c.Get(),
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
