package sshkeys

import (
	"strconv"

	"github.com/cherryservers/cherrygo/v4"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Command) List() *cobra.Command {
	var projectID int
	sshListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves SSH keys.",
		Long:  "Retrieves SSH keys. If project ID is specified, will return all SSH keys assigned to a specific project.",
		Example: `  # List SSH keys:
  cherryctl ssh-key list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			ctx := cmd.Context()

			getOptions := c.GetOpts()
			if len(getOptions.Fields) == 0 {
				getOptions.Fields = []string{"ssh_key", "email"}
			}

			var sshKeys []cherrygo.SSHKey
			err := error(nil)
			if projectID != 0 {
				sshKeys, _, err = c.ProjectClient().ListSSHKeys(ctx, projectID, getOptions)
			} else {
				sshKeys, _, err = c.Client().List(ctx, getOptions)
			}

			if err != nil {
				return errors.Wrap(err, "Could not get SSH key list")
			}

			data := make([][]string, len(sshKeys))
			for i, o := range sshKeys {
				data[i] = []string{strconv.Itoa(o.ID), o.Label, o.User.Email, o.Fingerprint, o.Created}
			}
			header := []string{"ID", "Label", "User", "Fingerprint", "Created"}

			return c.Outputer().Output(sshKeys, header, &data)
		},
	}

	sshListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "Project to retrieve keys from.")

	return sshListCmd
}
