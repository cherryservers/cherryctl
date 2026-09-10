package fakes

import (
	"context"
	"errors"

	"github.com/cherryservers/cherrygo/v4"
)

var _ cherrygo.SSHKeysService = (*SSHKeyService)(nil)

type SSHKeyService struct {
	Calls []CallRecord
}

// Create implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Create(ctx context.Context, request *cherrygo.CreateSSHKey) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}

// Delete implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Delete(ctx context.Context, sshKeyID int) (*cherrygo.Response, error) {
	return nil, errors.New("not implemented")
}

// Get implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Get(ctx context.Context, sshKeyID int, opts *cherrygo.GetOptions) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}

// List implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) List(ctx context.Context, opts *cherrygo.GetOptions) ([]cherrygo.SSHKey, *cherrygo.Response, error) {
	return nil, nil, errors.New("not implemented")
}

// Update implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Update(ctx context.Context, sshKeyID int, request *cherrygo.UpdateSSHKey) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}
