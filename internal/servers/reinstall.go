package servers

import (
	"fmt"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Reinstall() *cobra.Command {
	var (
		serverID        int
		hostname        string
		image           string
		password        string
		sshKeys         []string
		osPartitionSize int
	)

	reinstallServerCmd := &cobra.Command{
		Use:   `reinstall -i <server_id> --hostname --image <image_slug> --password <password> [--ssh-keys <ssh_key_ids>] [--os-partition-size <size>]`,
		Short: "Reinstall a server.",
		Long:  "Reinstall the specified server.",
		Example: `  # Reinstall the specified server:
  cherryctl server reinstall -i 12345 -h staging-server-1 --image ubuntu_20_04 --password G1h2e_39Q9oT`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			request := &cherrygo.ReinstallServerFields{
				Image:           image,
				Hostname:        hostname,
				Password:        password,
				SSHKeys:         sshKeys,
				OSPartitionSize: osPartitionSize,
			}

			_, _, err := c.Service.Reinstall(serverID, request)
			if err != nil {
				return errors.Wrap(err, "Could not reinstall Server")
			}

			fmt.Println("Server", serverID, "reinstall has been started.")
			return nil
		},
	}

	reinstallServerCmd.Flags().IntVarP(&serverID, "server-id", "i", 0, "The ID of a server.")
	reinstallServerCmd.Flags().StringVarP(&hostname, "hostname", "", "", "Hostname.")
	reinstallServerCmd.Flags().StringVarP(&image, "image", "", "", "Operating system slug for the server.")
	reinstallServerCmd.Flags().StringVarP(&password, "password", "", "", "Server password.")
	reinstallServerCmd.Flags().StringSliceVarP(&sshKeys, "ssh-keys", "", []string{}, "Comma separated list of SSH key IDs to be embed in the Server.")
	reinstallServerCmd.Flags().IntVarP(&osPartitionSize, "os-partition-size", "", 0, "OS partition size in GB.")

	_ = reinstallServerCmd.MarkFlagRequired("server-id")
	_ = reinstallServerCmd.MarkFlagRequired("hostname")
	_ = reinstallServerCmd.MarkFlagRequired("image")
	_ = reinstallServerCmd.MarkFlagRequired("password")

	return reinstallServerCmd
}
