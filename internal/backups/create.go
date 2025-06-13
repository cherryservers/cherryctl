package backups

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		projectID      int
		serverID       int
		serverHostname string
		backupPlan     string
		region         string
		sshKey         string
	)
	backupCreateCmd := &cobra.Command{
		Use:   `create {--server-id <id> | --server-hostname <hostname>} --plan <backup_plan_slug> --region <region_slug> [-p <project_id>] [--ssh-key <plain_ssh_key>]`,
		Short: "Create a backup storage.",
		Long:  "Create a backup storage for specified server.",
		Example: `  # Create backup storage with 100GB of space in the LT-Siauliai location for the server with hostname "delicate-zebra":
  cherryctl backup create --server-hostname delicate-zebra --plan backup_100 --region LT-Siauliai --project-id 123`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if serverHostname == "" && serverID == 0 {
				return fmt.Errorf("either server-id or server-hostname should be set")
			}

			request := &cherrygo.CreateBackup{
				ServerID:       serverID,
				BackupPlanSlug: backupPlan,
				RegionSlug:     region,
				SSHKey:         sshKey,
			}

			if serverHostname != "" {
				srvID, err := utils.ServerHostnameToID(serverHostname, projectID, c.ServerService)
				if err != nil {
					return errors.Wrap(err, "Could not get a Server")
				}
				request.ServerID = srvID
			}

			o, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not create a backup storage")
			}

			header := []string{"ID", "Status", "Attached to", "Size GB", "Used GB", "Region", "Price"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), o.Status, o.AttachedTo.Hostname, strconv.Itoa(o.SizeGigabytes), strconv.Itoa(o.UsedGigabytes), o.Region.Slug, fmt.Sprint(o.Pricing.Price)}

			return c.Out.Output(o, header, &data)
		},
	}

	backupCreateCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	backupCreateCmd.Flags().IntVarP(&serverID, "server-id", "s", 0, "The server's ID.")
	backupCreateCmd.Flags().StringVarP(&serverHostname, "server-hostname", "", "", "The Hostname of a server.")
	backupCreateCmd.Flags().StringVarP(&region, "region", "", "", "Slug of the region.")
	backupCreateCmd.Flags().StringVarP(&backupPlan, "plan", "", "", "Backup storage plan slug.")
	backupCreateCmd.Flags().StringVarP(&sshKey, "ssh-key", "", "", "Plain SSH key will be stored in backup service.")

	backupCreateCmd.MarkFlagRequired("region")
	backupCreateCmd.MarkFlagRequired("plan")

	backupCreateCmd.MarkFlagsMutuallyExclusive("server-id", "server-hostname")

	return backupCreateCmd
}
