package teams

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var teamID int
	var force bool
	deleteTeamCmd := &cobra.Command{
		Use: `delete ID -t <team_id>`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				tID, err := strconv.Atoi(args[0])
				if err == nil {
					teamID = tID
				}
			}
			return nil
		},
		Short: "Delete a team.",
		Long:  "Deletes the specified team with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified team:
  cherryctl team delete 12345
  >
  ✔ Are you sure you want to delete team 12345: y
  		
  # Deletes a team, skipping confirmation:
  cherryctl team delete -f -t 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete team %d? All asociated resources (servers, IP addresses, storages, etc.) will be terminated also. ", teamID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			_, err := c.Service.Delete(teamID)
			if err != nil {
				return errors.Wrap(err, "Could not delete a Team")
			}

			fmt.Println("Team", teamID, "successfully deleted.")
			return nil
		},
	}

	deleteTeamCmd.Flags().IntVarP(&teamID, "team-id", "t", 0, "The ID of a team.")
	deleteTeamCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the tean deletion.")

	return deleteTeamCmd
}
