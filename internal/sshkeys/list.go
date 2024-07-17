package sshkeys

import (
	"github.com/cherryservers/cherrygo/v3"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	sshListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves ssh-keys.",
		Long:  "Retrieves ssh-keys. If the project ID is specified, will return all SSH keys assigned to a specific project.",
		Example: `  # List of ssh-keys:
  cherryctl ssh-key list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"ssh_key", "email"}

			var sshKeys []cherrygo.SSHKey
			err := error(nil)
			if projectID != 0 {
				sshKeys, _, err = c.ProjectsService.ListSSHKeys(projectID, getOptions)
			} else {
				sshKeys, _, err = c.Service.List(getOptions)
			}

			if err != nil {
				return errors.Wrap(err, "Could not get ssh-keys list")
			}

			data := make([][]string, len(sshKeys))
			for i, o := range sshKeys {
				data[i] = []string{strconv.Itoa(o.ID), o.Label, o.User.Email, o.Fingerprint, o.Created}
			}
			header := []string{"ID", "Label", "User", "Fingerprint", "Created"}

			return c.Out.Output(sshKeys, header, &data)
		},
	}

	sshListCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")

	return sshListCmd
}
