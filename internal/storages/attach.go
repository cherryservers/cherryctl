package storages

import (
	"fmt"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Attach() *cobra.Command {
	var (
		storageID      int
		serverID       int
		serverHostname string
		projectID      int
	)
	storageAttachCmd := &cobra.Command{
		Use:   `attach -i <storage_id> {--server-id | --server-hostname} [-p <project_id>]`,
		Short: "Attach storage volume to a specified server.",
		Long:  "Attach storage volume to a specified server.",
		Example: `  # Attach storage to specified server:
  cherryctl storage attach -i 12345 -s 12345`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if serverHostname == "" && serverID == 0 {
				return fmt.Errorf("either server-id or server-hostname should be set")
			}

			request := &cherrygo.AttachTo{
				StorageID: storageID,
			}

			if serverHostname != "" {
				srvID, err := utils.ServerHostnameToID(serverHostname, projectID, c.ServerService)
				if err != nil {
					return errors.Wrap(err, "Could not get a Server")
				}
				request.AttachTo = srvID
			}

			resp, _, err := c.Service.Attach(request)
			if err != nil {
				return errors.Wrap(err, "Could not atach storage to a server")
			}

			fmt.Println("Storage", storageID, "successfully attached to", resp.AttachedTo.Hostname)
			return nil
		},
	}

	storageAttachCmd.Flags().IntVarP(&storageID, "storage-id", "i", 0, "The storage's ID.")
	storageAttachCmd.Flags().IntVarP(&serverID, "server-id", "s", 0, "The server's ID.")
	storageAttachCmd.Flags().StringVarP(&serverHostname, "server-hostname", "", "", "The Hostname of a server.")
	storageAttachCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")

	storageAttachCmd.MarkFlagsMutuallyExclusive("server-id", "server-hostname")

	storageAttachCmd.MarkFlagRequired("storage-id")

	return storageAttachCmd
}
