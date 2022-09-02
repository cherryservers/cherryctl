package sshkeys

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var projectID int
	sshListCmd := &cobra.Command{
		Use:   `list [-p <project_id>]`,
		Short: "Retrieves project members ssh-keys details.",
		Long:  "Retrieves project members ssh-keys details.",
		Example: `  # List of project ssh-keys:
  cherryctl ssh-key list -i 12345`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOptions := c.Servicer.GetOptions()
			getOptions.Fields = []string{"ssh_key", "email"}
			sshKeys, _, err := c.ProjectsService.ListSSHKeys(projectID, getOptions)
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
	sshListCmd.MarkFlagRequired("project-id")

	return sshListCmd
}
