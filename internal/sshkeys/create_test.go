package sshkeys_test

import (
	"errors"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"create", "--key", "a", "--label", "a"}, fakeDeps)

	require.NoError(t, cmd.Execute())

	require.Len(t, fakeDeps.svc.Calls, 1)
	fakeDeps.svc.Calls[0].AssertMethod(t, "Create")
	fakeDeps.svc.Calls[0].AssertParams(t, []any{t.Context(), &cherrygo.CreateSSHKey{
		Label: "a",
		Key:   "a",
	}}...)

	require.Len(t, fakeDeps.out.Calls, 1)
	wantTh := []string{"ID", "Label", "Fingerprint", "Created"}
	wantTd := [][]string{{"1", "test-label", "test-fingerprint", "test-created"}}
	fakeDeps.out.Calls[0].Assert(t, fakes.SSHKey(), wantTh, wantTd)
}

func TestCreateInputErrors(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "with positional args",
			args:    []string{"create", "--key", "a", "--label", "a", "b"},
			wantErr: "unknown command \"b\"",
		},
		{
			name:    "without key",
			args:    []string{"create", "--label", "a"},
			wantErr: "required flag",
		},
		{
			name:    "without label",
			args:    []string{"create", "--key", "a"},
			wantErr: "required flag",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDeps := newFakeDeps()
			cmd := setupCommand(t.Context(), tc.args, fakeDeps)

			require.ErrorContains(t, cmd.Execute(), tc.wantErr)
			assert.Empty(t, fakeDeps.svc.Calls)
			assert.Empty(t, fakeDeps.out.Calls)
		})
	}
}

func TestCreateClientError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.svc.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"create", "--key", "a", "--label", "a"}, fakeDeps)

	require.ErrorContains(t, cmd.Execute(), "Could not create SSH key")
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Empty(t, fakeDeps.out.Calls)
}

func TestCreateOutputError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.out.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"create", "--key", "a", "--label", "a"}, fakeDeps)

	require.ErrorContains(t, cmd.Execute(), "test")
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Len(t, fakeDeps.out.Calls, 1)
}
