package fakes

import (
	"context"
	"errors"

	"github.com/cherryservers/cherrygo/v4"
)

var _ cherrygo.SSHKeysService = (*SSHKeyService)(nil)

type SSHKeyService struct {
	Calls []CallRecord

	// Err is the error returned with every implemented method.
	// Useful for faking client errors.
	Err error
}

// Create implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Create(_ context.Context, _ *cherrygo.CreateSSHKey) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}

// Delete implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Delete(_ context.Context, _ int) (*cherrygo.Response, error) {
	return nil, errors.New("not implemented")
}

// Get implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Get(_ context.Context, _ int, _ *cherrygo.GetOptions) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}

// List implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) List(_ context.Context, _ *cherrygo.GetOptions) ([]cherrygo.SSHKey, *cherrygo.Response, error) {
	return nil, nil, errors.New("not implemented")
}

// Update implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Update(ctx context.Context, id int, request *cherrygo.UpdateSSHKey) (cherrygo.SSHKey, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "Update", params: []any{ctx, id, request}})
	return SSHKey(), nil, s.Err
}

// Key is the SSH key returned by the fake methods.
func SSHKey() cherrygo.SSHKey {
	return cherrygo.SSHKey{
		ID:          1,
		Label:       "test-label",
		Fingerprint: "test-fingerprint",
		User:        cherrygo.User{Email: "test-email"},
		Created:     "test-created",
	}
}
