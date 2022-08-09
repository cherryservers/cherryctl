package sshkeys

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer
	Service         cherrygo.SSHKeysService
	ProjectsService cherrygo.ProjectsService
	Out             outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `ssh-key`,
		Aliases: []string{"sshkey", "sshkeys", "ssh-keys"},
		Short:   "Ssh-key operations.",
		Long:    "Ssh-key operations: get, list, update, delete.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = c.Servicer.API(cmd).SSHKeys
			c.ProjectsService = c.Servicer.API(cmd).Projects
		},
	}

	cmd.AddCommand(
		c.Get(),
		c.List(),
		c.Create(),
		c.Update(),
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
