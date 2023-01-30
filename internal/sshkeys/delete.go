package sshkeys

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var sshKeyID int
	var force bool
	deleteSSHkeyCmd := &cobra.Command{
		Use:   `delete -i <ssh_key_id> [-f]`,
		Short: "Deletes an SSH key.",
		Long:  "Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.",
		Example: `  # Deletes an SSH key, with confirmation:
  cherryctl shh-key delete -i 12345
  >
  ✔ Are you sure you want to delete SSH key 12345: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl shh-key delete -f -i 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete SSH key %d: ", sshKeyID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			_, _, err := c.Service.Delete(sshKeyID)
			if err != nil {
				return errors.Wrap(err, "Could not delete SSH key")
			}

			fmt.Println("SSH key", sshKeyID, "successfully deleted.")
			return nil
		},
	}

	deleteSSHkeyCmd.Flags().IntVarP(&sshKeyID, "ssh-key-id", "i", 0, "ID of the SSH key.")
	deleteSSHkeyCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the SSH key deletion.")

	_ = deleteSSHkeyCmd.MarkFlagRequired("ssh-key-id")

	return deleteSSHkeyCmd
}
