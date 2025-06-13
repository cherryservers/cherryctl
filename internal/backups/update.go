package backups

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		backupID int
		plan     string
		password string
		sshKey   string
	)
	backupUpdateCmd := &cobra.Command{
		Use:   `update <backup_ID> [--password <plain_text>] [--ssh-key <plain_ssh_key>]`,
		Args:  cobra.ExactArgs(1),
		Short: "Update a backup storage.",
		Long:  "Update the backup user password or SSH key. Passwords are used in the FTP and SMB protocols, while SSH keys are used in BORG.",
		Example: `  # Update backup storage password and SSH key:
  cherryctl backup update 12345 --password somePassword --ssh-key  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6ec8eT..."
  
  # Update backup storage user password:
  cherryctl backup update 12345 --password somePassword`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if backID, err := strconv.Atoi(args[0]); err == nil {
				backupID = backID
			}
			request := &cherrygo.UpdateBackupStorage{
				BackupStorageID: backupID,
				BackupPlanSlug:  plan,
				Password:        password,
				SSHKey:          sshKey,
			}

			o, _, err := c.Service.Update(request)
			if err != nil {
				return errors.Wrap(err, "Could not update backup storage")
			}

			header := []string{"ID", "Status", "Attached to", "Size GB", "Used GB", "Region", "Price"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Status, o.AttachedTo.Hostname, strconv.Itoa(o.SizeGigabytes), strconv.Itoa(o.UsedGigabytes), o.Region.Slug, fmt.Sprint(o.Pricing.Price)}

			return c.Out.Output(o, header, &data)
		},
	}

	backupUpdateCmd.Flags().StringVarP(&password, "password", "", "", "Backup storage user password.")
	backupUpdateCmd.Flags().StringVarP(&sshKey, "ssh-key", "", "", "Plain SSH key which will be stored in the backup service.")

	return backupUpdateCmd
}
