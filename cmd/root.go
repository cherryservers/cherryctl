/*
Copyright © 2022 Cherry Severs <support@cherryservers.com>
*/
package cmd

import (
	"github.com/cherryservers/cherryctl/internal/backups"
	root "github.com/cherryservers/cherryctl/internal/cli"
	"github.com/cherryservers/cherryctl/internal/docs"
	"github.com/cherryservers/cherryctl/internal/images"
	initPck "github.com/cherryservers/cherryctl/internal/init"
	"github.com/cherryservers/cherryctl/internal/ips"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherryctl/internal/plans"
	"github.com/cherryservers/cherryctl/internal/projects"
	"github.com/cherryservers/cherryctl/internal/regions"
	"github.com/cherryservers/cherryctl/internal/servers"
	"github.com/cherryservers/cherryctl/internal/sshkeys"
	"github.com/cherryservers/cherryctl/internal/storages"
	"github.com/cherryservers/cherryctl/internal/teams"
	"github.com/cherryservers/cherryctl/internal/users"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
)

var Version string = "dev"

type Cli struct {
	MainCmd  *cobra.Command
	Outputer outputs.Outputer
}

func NewCli() *Cli {
	cli := &Cli{
		Outputer: &outputs.Standard{},
	}

	rootClient := root.NewClient(Version)
	rootCmd := rootClient.NewCommand()
	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd

	cli.RegisterCommands(rootClient)

	cobra.OnInitialize(
		func() {
			cli.Outputer.SetFormat(rootClient.Format())
		},
	)
	return cli
}

type sharedDeps struct {
	out    outputs.Outputer
	client *root.Client
}

func (d sharedDeps) GetOpts() *cherrygo.GetOptions {
	return d.client.GetOptions()
}

func (d sharedDeps) Outputer() outputs.Outputer {
	return d.out
}

type planDeps struct {
	client *root.Client
	sharedDeps
}

func (d *planDeps) Client() cherrygo.PlansService {
	// The API method doesn't actually use the command for anything, so we can pass nil.
	// Should refactor it out at some point.
	return d.client.API(nil).Plans
}

type serverDeps struct {
	client *root.Client
	sharedDeps
}

func (d *serverDeps) Client() cherrygo.ServersService {
	return d.client.API(nil).Servers
}

func (cli *Cli) RegisterCommands(client *root.Client) {
	shared := sharedDeps{
		out:    cli.Outputer,
		client: client,
	}
	cli.MainCmd.AddCommand(
		docs.NewCommand(),

		initPck.NewClient(client).NewCommand(),
		servers.NewCommand(&serverDeps{client: client, sharedDeps: shared}).CobraCommand(),
		ips.NewClient(client, cli.Outputer).NewCommand(),
		storages.NewClient(client, cli.Outputer).NewCommand(),
		backups.NewClient(client, cli.Outputer).NewCommand(),
		regions.NewClient(client, cli.Outputer).NewCommand(),
		// We don't have the dependencies initialized yet, as that's done
		// on pre-execution, so we need an injector interface.
		plans.NewCommand(&planDeps{client: client, sharedDeps: shared}).CobraCommand(),
		projects.NewClient(client, cli.Outputer).NewCommand(),
		teams.NewClient(client, cli.Outputer).NewCommand(),
		sshkeys.NewClient(client, cli.Outputer).NewCommand(),
		images.NewClient(client, cli.Outputer).NewCommand(),
		users.NewClient(client, cli.Outputer).NewCommand(),
	)
}
