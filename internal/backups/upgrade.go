package backups

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Upgrade() *cobra.Command {
	var (
		backupID int
		plan     string
	)
	backupUpgradeCmd := &cobra.Command{
		Use:   `upgrade <backup_ID> --plan <backup_plan_slug>`,
		Args:  cobra.ExactArgs(1),
		Short: "Upgrade a backup storage plan.",
		Long:  "Upgrade a backup storage plan to increase it's storage size. ATTENTION! Upgrade can be done once per backup plan.",
		Example: `  # Upgrade backup storage size to 1000 gigabytes:
  cherryctl backup upgrade 12345 --plan backup_1000`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if backID, err := strconv.Atoi(args[0]); err == nil {
				backupID = backID
			}
			request := &cherrygo.UpdateBackupStorage{
				BackupStorageID: backupID,
				BackupPlanSlug:  plan,
			}

			o, _, err := c.Service.Update(request)
			if err != nil {
				return errors.Wrap(err, "Could not upgrade backup storage")
			}

			header := []string{"ID", "Status", "Attached to", "Size GB", "Used GB", "Region", "Price"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Status, o.AttachedTo.Hostname, strconv.Itoa(o.SizeGigabytes), strconv.Itoa(o.UsedGigabytes), o.Region.Slug, fmt.Sprint(o.Pricing.Price)}

			return c.Out.Output(o, header, &data)
		},
	}

	backupUpgradeCmd.Flags().StringVarP(&plan, "plan", "", "", "Backup storage plan slug.")

	return backupUpgradeCmd
}
