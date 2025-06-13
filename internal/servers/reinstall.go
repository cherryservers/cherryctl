package servers

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Reinstall() *cobra.Command {
	var (
		hostname        string
		image           string
		password        string
		sshKeys         []string
		userDataFile    string
		userdata        string
		osPartitionSize int
	)

	reinstallServerCmd := &cobra.Command{
		Use:   `reinstall ID --hostname <hostname> --image <image_slug> --password <password> [--ssh-keys <ssh_key_ids>] [--os-partition-size <size>] [--userdata-file <filepath>]`,
		Args:  cobra.ExactArgs(1),
		Short: "Reinstall a server.",
		Long:  "Reinstall the specified server.",
		Example: `  # Reinstall the specified server:
  cherryctl server reinstall 12345 --hostname staging-server-1 --image ubuntu_24_04 --password G1h2e_39Q9oT`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if userDataFile != "" {
				userdataRaw, readErr := os.ReadFile(userDataFile)
				if readErr != nil {
					return errors.Wrap(readErr, "Could not read userdata-file")
				}
				userdata = base64.StdEncoding.EncodeToString(userdataRaw)
			}

			request := &cherrygo.ReinstallServerFields{
				Image:           image,
				Hostname:        hostname,
				Password:        password,
				SSHKeys:         sshKeys,
				UserData:        userdata,
				OSPartitionSize: osPartitionSize,
			}

			if serverID, err := strconv.Atoi(args[0]); err == nil {
				_, _, err := c.Service.Reinstall(serverID, request)
				if err != nil {
					return errors.Wrap(err, "Could not reinstall a Server.")
				}

				fmt.Println("Server", serverID, "reinstall has been started.")
				return nil
			}

			fmt.Println("Server with ID %s was not found.", args[0])
			return nil
		},
	}

	reinstallServerCmd.Flags().StringVarP(&hostname, "hostname", "", "", "Hostname.")
	reinstallServerCmd.Flags().StringVarP(&image, "image", "", "", "Operating system slug for the server.")
	reinstallServerCmd.Flags().StringVarP(&password, "password", "", "", "Server password.")
	reinstallServerCmd.Flags().StringSliceVarP(&sshKeys, "ssh-keys", "", []string{}, "Comma separated list of SSH key IDs to be embed in the Server.")
	reinstallServerCmd.Flags().IntVarP(&osPartitionSize, "os-partition-size", "", 0, "OS partition size in GB.")
	reinstallServerCmd.Flags().StringVarP(&userDataFile, "userdata-file", "", "", "Path to a userdata file for server initialization.")

	_ = reinstallServerCmd.MarkFlagRequired("hostname")
	_ = reinstallServerCmd.MarkFlagRequired("image")
	_ = reinstallServerCmd.MarkFlagRequired("password")

	return reinstallServerCmd
}
