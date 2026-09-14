package sshkeys_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const wantPromptMsg = "Are you sure you want to delete SSH key 1"

func TestDelete(t *testing.T) {
	cases := []struct {
		name               string
		args               []string
		wantPromptMessages []string
	}{
		{
			name:               "with confirmation",
			args:               []string{"delete", "--ssh-key-id", "1"},
			wantPromptMessages: []string{wantPromptMsg},
		},
		{
			name: "forced",
			args: []string{"delete", "--ssh-key-id", "1", "--force"},
		},
		{
			name: "shorthands",
			args: []string{"delete", "-i", "1", "-f"},
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

			assert.Equal(t, "SSH key 1 successfully deleted.\n", buf.String())
		})
	}
}

func TestDeleteNoDeletionWhenPromptDenied(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	cmd := setupCommand(t.Context(), []string{"delete", "-i", "1"}, fakeDeps)
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

	require.ErrorContains(t, cmd.Execute(), "not set")
	assert.Empty(t, buf)
	assert.Empty(t, fakeDeps.svc.Calls)
}

func TestDeleteErrorWhenPromptError(t *testing.T) {
	var buf bytes.Buffer

	fakeDeps := newFakeDeps()
	fakeDeps.prompt.Err = errors.New("test")
	cmd := setupCommand(t.Context(), []string{"delete", "-i", "1"}, fakeDeps)
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
	cmd := setupCommand(t.Context(), []string{"delete", "-i", "1"}, fakeDeps)
	cmd.SetOut(&buf)

	require.ErrorContains(t, cmd.Execute(), "Could not delete")
	assert.Empty(t, buf)
	assert.Len(t, fakeDeps.svc.Calls, 1)
}
