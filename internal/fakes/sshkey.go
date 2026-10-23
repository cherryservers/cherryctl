package fakes

import (
	"context"
	"errors"

	"github.com/cherryservers/cherrygo/v4"
)

var _ cherrygo.SSHKeysService = (*SSHKeyService)(nil)

type SSHKeyService struct {
	fakeService
}

// Create implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Create(_ context.Context, _ *cherrygo.CreateSSHKey) (cherrygo.SSHKey, *cherrygo.Response, error) {
	return cherrygo.SSHKey{}, nil, errors.New("not implemented")
}

// Delete implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Delete(ctx context.Context, id int) (*cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "Delete", params: []any{ctx, id}})
	return nil, s.Err
}

// Get implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) Get(ctx context.Context, id int, opts *cherrygo.GetOptions) (cherrygo.SSHKey, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "Get", params: []any{ctx, id, opts}})
	return SSHKey(), nil, s.Err
}

// List implements [cherrygo.SSHKeysService].
func (s *SSHKeyService) List(ctx context.Context, opts *cherrygo.GetOptions) ([]cherrygo.SSHKey, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "List", params: []any{ctx, opts}})
	return []cherrygo.SSHKey{SSHKey()}, nil, s.Err
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
