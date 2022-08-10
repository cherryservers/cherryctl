/*
Copyright © 2022 Cherry Severs <support@cherryservers.com>

*/
package cmd

import (
	"os"

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
	"github.com/spf13/cobra"
)

var (
	Version string = "dev"
)

const (
	apiTokenEnvVar = "CHERRY_AUTH_TOKEN"
	apiURL         = "https://api.cherryservers.com/v1/"
)

type Cli struct {
	MainCmd  *cobra.Command
	Outputer outputs.Outputer
}

func NewCli() *Cli {
	cli := &Cli{
		Outputer: &outputs.Standard{},
	}

	rootClient := root.NewClient(os.Getenv(apiTokenEnvVar), apiURL, Version)
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

func (cli *Cli) RegisterCommands(client *root.Client) {
	cli.MainCmd.AddCommand(
		docs.NewCommand(),

		initPck.NewClient(client).NewCommand(),
		servers.NewClient(client, cli.Outputer).NewCommand(),
		ips.NewClient(client, cli.Outputer).NewCommand(),
		storages.NewClient(client, cli.Outputer).NewCommand(),
		regions.NewClient(client, cli.Outputer).NewCommand(),
		plans.NewClient(client, cli.Outputer).NewCommand(),
		projects.NewClient(client, cli.Outputer).NewCommand(),
		teams.NewClient(client, cli.Outputer).NewCommand(),
		sshkeys.NewClient(client, cli.Outputer).NewCommand(),
		images.NewClient(client, cli.Outputer).NewCommand(),
		users.NewClient(client, cli.Outputer).NewCommand(),
	)
}
