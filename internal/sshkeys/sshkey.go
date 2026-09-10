package sshkeys

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
)

type Command struct {
	Deps
}

func (c *Command) CobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `ssh-key`,
		Aliases: []string{"sshkey", "sshkeys", "ssh-keys"},
		Short:   "Ssh-key operations.",
		Long:    "Ssh-key operations: get, list, update, delete.",
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

type Deps interface {
	Client() cherrygo.SSHKeysService
	ProjectClient() cherrygo.ProjectsService
	GetOpts() *cherrygo.GetOptions
	Outputer() outputs.Outputer
}

func NewCommand(dep Deps) *Command {
	return &Command{
		Deps: dep,
	}
}
