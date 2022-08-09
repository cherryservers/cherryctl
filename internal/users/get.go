package users

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var userID int
	userGetCmd := &cobra.Command{
		Use:   `get [-i <user_id>]`,
		Short: "Retrieves information about the current user or a specified user.",
		Long:  "Returns either information about the current user or information about a specified user. Specified user information is only available if that user shares a project with the current user.",
		Example: `  # Gets the current user's information:
		cherryctl user get
		
		# Returns information on user with ID 123:
		cherryctl user get -i 123`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var user cherrygo.User
			var err error

			if userID == 0 {
				user, _, err = c.Service.CurrentUser(c.Servicer.GetOptions())
				if err != nil {
					return errors.Wrap(err, "Could not get current User")
				}
			} else {
				user, _, err = c.Service.Get(userID, c.Servicer.GetOptions())
				if err != nil {
					return errors.Wrap(err, "Could not get Users")
				}
			}

			data := make([][]string, 1)

			data[0] = []string{strconv.Itoa(user.ID), fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email}
			header := []string{"ID", "Full Name", "Email"}

			return c.Out.Output(user, header, &data)
		},
	}

	userGetCmd.Flags().IntVarP(&userID, "user-id", "i", 0, "The ID of the user.")

	return userGetCmd
}
