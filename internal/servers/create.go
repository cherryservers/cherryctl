package servers

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		projectID       int
		hostname        string
		image           string
		sshKeys         []string
		osPartitionSize int
		region          string
		plan            string
		userDataFile    string
		userdata        string
		tags            []string
		spotInstance    bool
		ipAddresses     []string
		storageID       int
	)

	createServerCmd := &cobra.Command{
		Use:   `create -p <project_id> --plan <plan_slug> --region <region_slug> [--hostname <hostname>] [--image <image_slug>] [--ssh-keys <ssh_key_ids>] [--ip-addresses <ip_addresses_ids>] [--os-partition-size <size>] [--userdata-file <filepath>] [--tags] [--spot-instance] [--storage-id <storage_id>]`,
		Short: "Create a server.",
		Long:  "Create a server in specified project.",
		Example: `  # Provisions a E5-1620v4 server in the LT-Siauliai location running on Ubuntu 24.04:
  cherryctl server create -p <project_id> --plan e5_1620v4 --hostname staging-server-1 --image ubuntu_24_04 --region LT-Siauliai`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			tagsArr := make(map[string]string)

			if userDataFile != "" {
				userdataRaw, readErr := ioutil.ReadFile(userDataFile)
				if readErr != nil {
					return errors.Wrap(readErr, "Could not read userdata-file")
				}
				userdata = base64.StdEncoding.EncodeToString(userdataRaw)
			}

			for _, kv := range tags {
				var k, v string
				tokens := strings.SplitN(kv, "=", 2)
				k = strings.TrimSpace(tokens[0])
				if len(tokens) != 1 {
					v = strings.TrimSpace(tokens[1])
				}

				tagsArr[k] = v
			}

			request := &cherrygo.CreateServer{
				ProjectID:       projectID,
				Plan:            plan,
				Image:           image,
				Region:          region,
				Hostname:        hostname,
				IPAddresses:     ipAddresses,
				SSHKeys:         sshKeys,
				SpotInstance:    spotInstance,
				OSPartitionSize: osPartitionSize,
				UserData:        userdata,
				Tags:            &tagsArr,
				StorageID:       storageID,
			}

			s, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not provision a server")
			}

			header := []string{"ID", "Name", "Hostname", "Image", "State", "Region"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(s.ID), s.Name, s.Hostname, s.Image, s.State, s.Region.Name}

			return c.Out.Output(s, header, &data)
		},
	}

	createServerCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	createServerCmd.Flags().StringVarP(&plan, "plan", "", "", "Slug of the plan.")
	createServerCmd.Flags().StringVarP(&hostname, "hostname", "", "", "Server hostname.")
	createServerCmd.Flags().StringVarP(&image, "image", "", "", "Operating system slug for the server.")
	createServerCmd.Flags().StringSliceVarP(&sshKeys, "ssh-keys", "", []string{}, "Comma separated list of SSH key ID's to be embed in the Server.")
	createServerCmd.Flags().IntVarP(&osPartitionSize, "os-partition-size", "", 0, "OS partition size in GB.")
	createServerCmd.Flags().StringVarP(&region, "region", "", "", "Slug of the region where server will be provisioned.")
	createServerCmd.Flags().StringVarP(&userDataFile, "userdata-file", "", "", "Path to a userdata file for server initialization.")
	createServerCmd.Flags().BoolVarP(&spotInstance, "spot-instance", "", false, "Provisions the server as a spot instance.")
	createServerCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, `Tag or list of tags for the server: --tags="key=value,env=prod".`)
	createServerCmd.Flags().StringSliceVarP(&ipAddresses, "ip-addresses", "", []string{}, "Comma separated list of IP addresses ID's to be embed in the Server.")
	createServerCmd.Flags().IntVarP(&storageID, "storage-id", "", 0, "ID of the storage that will be attached to server.")

	_ = createServerCmd.MarkFlagRequired("project-id")
	_ = createServerCmd.MarkFlagRequired("plan")
	_ = createServerCmd.MarkFlagRequired("region")

	return createServerCmd
}
