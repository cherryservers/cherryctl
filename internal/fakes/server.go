package fakes

import (
	"context"

	"github.com/cherryservers/cherrygo/v4"
)

type ServerCreationFunc func(context.Context, *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error)

type ServerService struct {
	Calls  []CallRecord
	create ServerCreationFunc
}

// AllowBMCAccess implements [cherrygo.ServersService].
func (s *ServerService) AllowBMCAccess(ctx context.Context, serverID int, ip4 string) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

func (s *ServerService) SetCreate(f ServerCreationFunc) {
	s.create = f
}

// Create implements [cherrygo.ServersService].
func (s *ServerService) Create(ctx context.Context, request *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "Create", params: []any{request}})
	return s.create(ctx, request)
}

// Delete implements [cherrygo.ServersService].
func (s *ServerService) Delete(ctx context.Context, serverID int) (*cherrygo.Response, error) {
	panic("not implemented")
}

// EnterRescueMode implements [cherrygo.ServersService].
func (s *ServerService) EnterRescueMode(ctx context.Context, serverID int, fields *cherrygo.RescueServerFields) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// ExitRescueMode implements [cherrygo.ServersService].
func (s *ServerService) ExitRescueMode(ctx context.Context, serverID int) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// Get implements [cherrygo.ServersService].
func (s *ServerService) Get(ctx context.Context, serverID int, opts *cherrygo.GetOptions) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// List implements [cherrygo.ServersService].
func (s *ServerService) List(ctx context.Context, projectID int, opts *cherrygo.GetOptions) ([]cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// ListCycles implements [cherrygo.ServersService].
func (s *ServerService) ListCycles(ctx context.Context, opts *cherrygo.GetOptions) ([]cherrygo.ServerCycle, *cherrygo.Response, error) {
	panic("not implemented")
}

// ListSSHKeys implements [cherrygo.ServersService].
func (s *ServerService) ListSSHKeys(ctx context.Context, serverID int, opts *cherrygo.GetOptions) ([]cherrygo.SSHKey, *cherrygo.Response, error) {
	panic("not implemented")
}

// PowerOff implements [cherrygo.ServersService].
func (s *ServerService) PowerOff(ctx context.Context, serverID int) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// PowerOn implements [cherrygo.ServersService].
func (s *ServerService) PowerOn(ctx context.Context, serverID int) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// PowerState implements [cherrygo.ServersService].
func (s *ServerService) PowerState(ctx context.Context, serverID int) (cherrygo.PowerState, *cherrygo.Response, error) {
	panic("not implemented")
}

// Reboot implements [cherrygo.ServersService].
func (s *ServerService) Reboot(ctx context.Context, serverID int) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// Reinstall implements [cherrygo.ServersService].
func (s *ServerService) Reinstall(ctx context.Context, serverID int, fields *cherrygo.ReinstallServerFields) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// ResetBMCPassword implements [cherrygo.ServersService].
func (s *ServerService) ResetBMCPassword(ctx context.Context, serverID int) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// Update implements [cherrygo.ServersService].
func (s *ServerService) Update(ctx context.Context, serverID int, request *cherrygo.UpdateServer) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// Upgrade implements [cherrygo.ServersService].
func (s *ServerService) Upgrade(ctx context.Context, serverID int, plan string) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}

// WaitForStatus implements [cherrygo.ServersService].
func (s *ServerService) WaitForStatus(ctx context.Context, serverID int, status cherrygo.ServerStatus) (cherrygo.Server, *cherrygo.Response, error) {
	panic("not implemented")
}
