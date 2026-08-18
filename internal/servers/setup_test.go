package servers

import (
	"context"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type fakeDeps struct {
	svc  *fakes.ServerService
	out  *fakes.Outputer
	opts *cherrygo.GetOptions
}

func (fd fakeDeps) GetOpts() *cherrygo.GetOptions {
	return fd.opts
}

func (fd fakeDeps) Client() cherrygo.ServersService {
	return fd.svc
}

func (fd fakeDeps) Outputer() outputs.Outputer {
	return fd.out
}

func newTrue() *bool {
	b := true
	return &b
}

func setupCommand(svc *fakes.ServerService, out *fakes.Outputer) Command {
	dep := fakeDeps{svc: svc, out: out}
	return Command{Deps: dep}
}

func setupCobraCreateCommand(c Command, args []string) *cobra.Command {
	cmd := c.Create()
	cmd.SetArgs(args)
	cmd.SilenceUsage = true
	return cmd
}

func createOK(_ context.Context, _ *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error) {
	return cherrygo.Server{ID: 1}, nil, nil
}

func createErr(_ context.Context, _ *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error) {
	return cherrygo.Server{}, nil, errors.New("test-error")
}
