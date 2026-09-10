package sshkeys_test

import (
	"context"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherryctl/internal/sshkeys"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
)

type fakeDeps struct {
	svc *fakes.SSHKeyService
	out *fakes.Outputer
}

func (fd fakeDeps) GetOpts() *cherrygo.GetOptions {
	// Not required for updates.
	return nil
}

func (fd fakeDeps) Client() cherrygo.SSHKeysService {
	return fd.svc
}

func (fd fakeDeps) Outputer() outputs.Outputer {
	return fd.out
}

func (fd fakeDeps) ProjectClient() cherrygo.ProjectsService {
	// Not required for updates.
	return nil
}

func newFakeDeps() fakeDeps {
	return fakeDeps{
		svc: new(fakes.SSHKeyService),
		out: new(fakes.Outputer),
	}
}

func setupCommand(ctx context.Context, args []string, dep sshkeys.Deps) *cobra.Command {
	cmd := sshkeys.NewCommand(dep).CobraCommand()
	cmd.SetArgs(args)
	cmd.SetContext(ctx)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	return cmd
}
