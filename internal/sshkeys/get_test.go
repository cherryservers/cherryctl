package sshkeys_test

import (
	"errors"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name           string
		args           []string
		opts           *cherrygo.GetOptions
		wantClientOpts *cherrygo.GetOptions
	}{
		{
			name:           "only id",
			args:           []string{"get", "1"},
			opts:           &cherrygo.GetOptions{},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
		{
			name:           "preserve opt fields",
			args:           []string{"get", "1"},
			opts:           &cherrygo.GetOptions{Fields: []string{"test"}},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"test"}},
		},
		{
			name:           "overwrite empty opt fields",
			args:           []string{"get", "1"},
			opts:           &cherrygo.GetOptions{Fields: []string{}},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDeps := newFakeDeps()
			fakeDeps.opts = tc.opts
			cmd := setupCommand(t.Context(), tc.args, fakeDeps)

			require.NoError(t, cmd.Execute())

			require.Len(t, fakeDeps.svc.Calls, 1)
			fakeDeps.svc.Calls[0].AssertMethod(t, "Get")
			fakeDeps.svc.Calls[0].AssertParams(t, t.Context(), 1, tc.wantClientOpts)

			require.Len(t, fakeDeps.out.Calls, 1)
			wantTh := []string{"ID", "Label", "User", "Fingerprint", "Created"}
			wantTd := [][]string{{"1", "test-label", "test-email", "test-fingerprint", "test-created"}}
			fakeDeps.out.Calls[0].Assert(t, fakes.SSHKey(), wantTh, wantTd)
		})
	}
}

func TestGetReturnsErrorWhenNoID(t *testing.T) {
	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"get"}, fakeDeps)

	assert.ErrorContains(t, cmd.Execute(), "received 0")
	assert.Empty(t, fakeDeps.svc.Calls)
	assert.Empty(t, fakeDeps.projectSvc.Calls)
	assert.Empty(t, fakeDeps.out.Calls)
}

func TestGetReturnsErrorWhenInvalidID(t *testing.T) {
	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"get", "a"}, fakeDeps)

	assert.ErrorContains(t, cmd.Execute(), "invalid id")
	assert.Empty(t, fakeDeps.svc.Calls)
	assert.Empty(t, fakeDeps.projectSvc.Calls)
	assert.Empty(t, fakeDeps.out.Calls)
}

func TestGetReturnsErrorWhenClientError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.svc.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"get", "1"}, fakeDeps)

	assert.ErrorContains(t, cmd.Execute(), "Could not get SSH key")
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Empty(t, fakeDeps.projectSvc)
	assert.Empty(t, fakeDeps.out.Calls)
}

func TestGetReturnsErrorWhenOutputError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.out.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"get", "1"}, fakeDeps)

	assert.ErrorContains(t, cmd.Execute(), "test")
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Empty(t, fakeDeps.projectSvc)
	assert.Len(t, fakeDeps.out.Calls, 1)
}
