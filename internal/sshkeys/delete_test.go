package sshkeys_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	cases := []struct {
		name               string
		args               []string
		wantPromptMessages []string
		wantOut            string
	}{
		{
			name:               "with confirmation",
			args:               []string{"delete", "1"},
			wantPromptMessages: []string{"Are you sure you want to delete SSH key 1"},
			wantOut:            "SSH key 1 successfully deleted.\n",
		},
		{
			name:    "forced",
			args:    []string{"delete", "1", "--force"},
			wantOut: "SSH key 1 successfully deleted.\n",
		},
		{
			name:               "deprecated flag",
			args:               []string{"delete", "--ssh-key-id", "1"},
			wantPromptMessages: []string{"Are you sure you want to delete SSH key 1"},
			wantOut: "Flag --ssh-key-id has been deprecated, Pass the ID as an argument instead.\n" +
				"SSH key 1 successfully deleted.\n",
		},
		{
			name: "shorthands",
			args: []string{"delete", "-i", "1", "-f"},
			wantOut: "Flag --ssh-key-id has been deprecated, Pass the ID as an argument instead.\n" +
				"SSH key 1 successfully deleted.\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			fakeDeps := newFakeDeps()
			fakeDeps.prompt.OK = true
			cmd := setupCommand(t.Context(), tc.args, fakeDeps)
			cmd.SetOut(&buf)

			require.NoError(t, cmd.Execute())

			assert.Equal(t, tc.wantPromptMessages, fakeDeps.prompt.GotMessages)

			require.Len(t, fakeDeps.svc.Calls, 1)
			fakeDeps.svc.Calls[0].AssertMethod(t, "Delete")
			fakeDeps.svc.Calls[0].AssertParams(t, []any{t.Context(), 1}...)

			assert.Equal(t, tc.wantOut, buf.String())
		})
	}
}

func TestDeleteNoDeletionWhenPromptDenied(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"delete", "1"}, fakeDeps)
	cmd.SetOut(&buf)

	require.NoError(t, cmd.Execute())
	assert.Empty(t, buf)
	assert.Empty(t, fakeDeps.svc.Calls)
}

func TestDeleteErrorWhenNoID(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"delete"}, fakeDeps)
	cmd.SetOut(&buf)

	require.ErrorContains(t, cmd.Execute(), "no key id")
	assert.Empty(t, buf)
	assert.Empty(t, fakeDeps.svc.Calls)
}

func TestDeleteErrorWhenPromptError(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	fakeDeps.prompt.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"delete", "1"}, fakeDeps)
	cmd.SetOut(&buf)

	require.ErrorContains(t, cmd.Execute(), "test")
	assert.Empty(t, buf)
	assert.Empty(t, fakeDeps.svc.Calls)
}

func TestDeleteErrorWhenClientError(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	fakeDeps.prompt.OK = true
	fakeDeps.svc.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"delete", "1"}, fakeDeps)
	cmd.SetOut(&buf)

	require.ErrorContains(t, cmd.Execute(), "Could not delete")
	assert.Empty(t, buf)
	assert.Len(t, fakeDeps.svc.Calls, 1)
}

func TestDeleteIDInputErrors(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		wantErr string
		wantOut string
	}{
		{
			name:    "flag and positonal set",
			args:    []string{"delete", "--ssh-key-id", "1", "1"},
			wantErr: "both ssh-key-id flag and positional arg set",
			wantOut: "Flag --ssh-key-id has been deprecated, Pass the ID as an argument instead.\n",
		},
		{
			name:    "invalid positional arg",
			args:    []string{"delete", "a"},
			wantErr: "invalid id",
		},
		{
			name:    "no id set",
			args:    []string{"delete"},
			wantErr: "no key id",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			fakeDeps := newFakeDeps()
			cmd := setupCommand(t.Context(), tc.args, fakeDeps)
			cmd.SetOut(&buf)

			require.ErrorContains(t, cmd.Execute(), tc.wantErr)
			assert.Equal(t, tc.wantOut, buf.String())
			assert.Empty(t, fakeDeps.svc.Calls)
		})
	}
}
