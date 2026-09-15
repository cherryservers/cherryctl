package sshkeys_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	cases := []struct {
		name           string
		args           []string
		opts           *cherrygo.GetOptions
		wantClientOpts *cherrygo.GetOptions
	}{
		{
			name:           "no args",
			args:           []string{"list"},
			opts:           &cherrygo.GetOptions{},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
		{
			name:           "project id 0",
			args:           []string{"list", "--project-id", "0"},
			opts:           &cherrygo.GetOptions{},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
		{
			name:           "fields not overwritten when non-empty",
			args:           []string{"list"},
			opts:           &cherrygo.GetOptions{Fields: []string{"test"}},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"test"}},
		},
		{
			name:           "fields overwritten when empty",
			args:           []string{"list"},
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
			fakeDeps.svc.Calls[0].AssertMethod(t, "List")
			fakeDeps.svc.Calls[0].AssertParams(t, t.Context(), tc.wantClientOpts)

			require.Len(t, fakeDeps.out.Calls, 1)
			wantTh := []string{"ID", "Label", "User", "Fingerprint", "Created"}
			wantTd := [][]string{{
				"1", "test-label", "test-email", "test-fingerprint", "test-created",
			}}
			fakeDeps.out.Calls[0].Assert(t, []cherrygo.SSHKey{fakes.SSHKey()}, wantTh, wantTd)

			assert.Empty(t, fakeDeps.projectSvc.Calls)
		})
	}
}

func TestListByProject(t *testing.T) {
	cases := []struct {
		name           string
		flag           string
		opts           *cherrygo.GetOptions
		wantClientOpts *cherrygo.GetOptions
	}{
		{
			name:           "full flag",
			flag:           "--project-id",
			opts:           &cherrygo.GetOptions{},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
		{
			name:           "shorthand",
			flag:           "-p",
			opts:           &cherrygo.GetOptions{},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
		{
			name:           "fields not overwritten when non-empty",
			flag:           "-p",
			opts:           &cherrygo.GetOptions{Fields: []string{"test"}},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"test"}},
		},
		{
			name:           "fields overwritten when empty",
			flag:           "-p",
			opts:           &cherrygo.GetOptions{Fields: []string{}},
			wantClientOpts: &cherrygo.GetOptions{Fields: []string{"ssh_key", "email"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeDeps := newFakeDeps()
			fakeDeps.opts = tc.opts
			cmd := setupCommand(t.Context(), []string{"list", tc.flag, "1"}, fakeDeps)

			require.NoError(t, cmd.Execute())

			require.Len(t, fakeDeps.projectSvc.Calls, 1)
			fakeDeps.projectSvc.Calls[0].AssertMethod(t, "ListSSHKeys")
			fakeDeps.projectSvc.Calls[0].AssertParams(t, t.Context(), 1, tc.wantClientOpts)

			require.Len(t, fakeDeps.out.Calls, 1)
			wantTh := []string{"ID", "Label", "User", "Fingerprint", "Created"}
			wantTd := [][]string{{
				"1", "test-label", "test-email", "test-fingerprint", "test-created",
			}}
			fakeDeps.out.Calls[0].Assert(t, []cherrygo.SSHKey{fakes.SSHKey()}, wantTh, wantTd)

			assert.Empty(t, fakeDeps.svc.Calls)
		})
	}
}

func TestListReturnsErrorWhenClientError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.svc.Err = errors.New("")
	cmd := setupCommand(t.Context(), []string{"list"}, fakeDeps)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("Could not get SSH key list"), err.Error())
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Empty(t, fakeDeps.out.Calls)
	assert.Empty(t, fakeDeps.projectSvc.Calls)
}

func TestListReturnsErrorWhenProjectClientError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.projectSvc.Err = errors.New("")
	cmd := setupCommand(t.Context(), []string{"list", "--project-id", "1"}, fakeDeps)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("Could not get SSH key list"), err.Error())
	assert.Len(t, fakeDeps.projectSvc.Calls, 1)
	assert.Empty(t, fakeDeps.out.Calls)
	assert.Empty(t, fakeDeps.svc.Calls)
}

func TestListReturnsErrorWhenOutputError(t *testing.T) {
	fakeDeps := newFakeDeps()
	fakeDeps.out.Err = errors.New("")
	cmd := setupCommand(t.Context(), []string{"list"}, fakeDeps)

	err := cmd.Execute()
	require.Error(t, err)
	assert.Len(t, fakeDeps.svc.Calls, 1)
	assert.Len(t, fakeDeps.out.Calls, 1)
	assert.Len(t, fakeDeps.projectSvc.Calls, 0)
}
