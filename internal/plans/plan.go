package plans

import (
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
)

type Command struct {
	Deps
}

func (c *Command) CobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `plan`,
		Aliases: []string{"plans"},
		Short:   "Plan operations.",
		Long:    "Plan operations: get, list.",
	}

	cmd.AddCommand(
		c.list(),
	)

	return cmd
}

type Deps interface {
	Client() cherrygo.PlansService
	GetOpts() *cherrygo.GetOptions
	Outputer() outputs.Outputer
}

func NewCommand(dep Deps) *Command {
	return &Command{
		Deps: dep,
	}
}
