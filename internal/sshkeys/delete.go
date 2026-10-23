package sshkeys

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Command) Delete() *cobra.Command {
	var sshKeyID int
	var force bool
	deleteSSHkeyCmd := &cobra.Command{
		Use:   `delete ID [-f]`,
		Args:  cobra.MaximumNArgs(1),
		Short: "Deletes an SSH key.",
		Long:  "Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.",
		Example: `  # Deletes an SSH key, with confirmation:
  cherryctl shh-key delete 12345
  >
  ✔ Are you sure you want to delete SSH key 12345: y
  		
  # Deletes an SSH key, skipping confirmation:
  cherryctl shh-key delete 12345 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			cmd.SilenceUsage = true
			ctx := cmd.Context()

			if len(args) > 0 {
				if sshKeyID != 0 {
					return errors.New("both ssh-key-id flag and positional arg set")
				}
				sshKeyID, err = strconv.Atoi(args[0])
				if err != nil {
					return fmt.Errorf("invalid id: %w", err)
				}
			}
			if sshKeyID == 0 {
				return errors.New("no key id")
			}

			if !force {
				ok, err := c.PromptConfirmation(
					fmt.Sprintf("Are you sure you want to delete SSH key %d", sshKeyID),
				)
				if !ok || err != nil {
					return err
				}
			}
			_, err = c.Client().Delete(ctx, sshKeyID)
			if err != nil {
				return errors.Wrap(err, "Could not delete SSH key")
			}

			cmd.Println("SSH key", sshKeyID, "successfully deleted.")
			return nil
		},
	}

	deleteSSHkeyCmd.Flags().IntVarP(&sshKeyID, "ssh-key-id", "i", 0, "ID of the SSH key.")
	deleteSSHkeyCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation.")

	_ = deleteSSHkeyCmd.Flags().MarkDeprecated("ssh-key-id", "Pass the ID as an argument instead.")

	return deleteSSHkeyCmd
}
