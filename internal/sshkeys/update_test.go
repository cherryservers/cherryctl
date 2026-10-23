package sshkeys_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		name             string
		args             []string
		wantClientParams []any
	}{
		{
			name:             "only id",
			args:             []string{"update", "1"},
			wantClientParams: []any{ctx, 1, &cherrygo.UpdateSSHKey{}},
		},
		{
			name:             "empty label and key",
			args:             []string{"update", "1", "--label", "", "--key", ""},
			wantClientParams: []any{ctx, 1, &cherrygo.UpdateSSHKey{}},
		},
		{
			name: "label and key",
			args: []string{"update", "1", "--label", "test-label", "--key", "test-key"},
			wantClientParams: []any{ctx, 1, &cherrygo.UpdateSSHKey{
				Label: utils.ToPtr("test-label"),
				Key:   utils.ToPtr("test-key"),
			}},
		},
		{
			name:             "with deprecated id flag",
			args:             []string{"update", "1", "--ssh-key-id", "2"},
			wantClientParams: []any{ctx, 1, &cherrygo.UpdateSSHKey{}},
		},
		{
			name:             "with deprecated id flag shorthand",
			args:             []string{"update", "1", "-i", "2"},
			wantClientParams: []any{ctx, 1, &cherrygo.UpdateSSHKey{}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDeps := newFakeDeps()
			cmd := setupCommand(ctx, tc.args, fakeDeps)

			err := cmd.Execute()
			require.NoError(t, err)

			require.Len(t, fakeDeps.svc.Calls, 1)
			fakeDeps.svc.Calls[0].AssertMethod(t, "Update")
			fakeDeps.svc.Calls[0].AssertParams(t, tc.wantClientParams...)

			require.Len(t, fakeDeps.out.Calls, 1)
			wantTh := []string{"ID", "Label", "Fingerprint", "Created"}
			wantTd := [][]string{{"1", "test-label", "test-fingerprint", "test-created"}}
			fakeDeps.out.Calls[0].Assert(t, fakes.SSHKey(), wantTh, wantTd)
		})
	}
}

func TestUpdateReturnsErrorWhenInvalidID(t *testing.T) {
	fakes := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"update", "a"}, fakes)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("invalid id"), err.Error())
	assert.Empty(t, fakes.svc.Calls)
	assert.Empty(t, fakes.out.Calls)
}

func TestUpdateReturnsErrorWhenNoID(t *testing.T) {
	fakes := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"update"}, fakes)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("received 0"), err.Error())
	assert.Empty(t, fakes.svc.Calls)
	assert.Empty(t, fakes.out.Calls)
}

func TestUpdateReturnsErrorWhenClientError(t *testing.T) {
	fakes := newFakeDeps()
	fakes.svc.Err = errors.New("")
	cmd := setupCommand(t.Context(), []string{"update", "1"}, fakes)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("Could not update SSH key"), err.Error())
	require.Len(t, fakes.svc.Calls, 1)
	require.Empty(t, fakes.out.Calls)
}

func TestUpdateReturnsErrorWhenOutputError(t *testing.T) {
	fakes := newFakeDeps()
	fakes.out.Err = errors.New("")
	cmd := setupCommand(t.Context(), []string{"update", "1"}, fakes)

	err := cmd.Execute()
	require.Error(t, err)
	require.Len(t, fakes.svc.Calls, 1)
	require.Len(t, fakes.out.Calls, 1)
}
