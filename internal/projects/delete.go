package projects

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var projectID int
	var force bool
	deleteProjectCmd := &cobra.Command{
		Use:   `delete -p <project_id>`,
		Short: "Delete a project.",
		Long:  "Deletes the specified project with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified project:
  cherryctl project delete -p 12345
  >
  ✔ Are you sure you want to delete project 12345: y
  		
  # Deletes a project, skipping confirmation:
  cherryctl project delete -f -p 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete project %d? All asociated resources (servers, IP addresses, storages, etc.) will be terminated also. ", projectID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			_, err := c.Service.Delete(projectID)
			if err != nil {
				return errors.Wrap(err, "Could not delete Project")
			}

			fmt.Println("Project", projectID, "successfully deleted.")
			return nil
		},
	}

	deleteProjectCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The ID of a project.")
	deleteProjectCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the project deletion.")

	_ = deleteProjectCmd.MarkFlagRequired("project-id")

	return deleteProjectCmd
}
