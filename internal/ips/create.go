package ips

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		projectID      int
		region         string
		ptrRecord      string
		aRecord        string
		targetIP       string
		targetID       int
		target         string
		targetHostname string
		tags           []string
	)
	ipCreateCmd := &cobra.Command{
		Use:   `create [-p <project_id>] --region <region_slug> [--target-hostname | --target-id | --target-ip-id] [--ptr-record] [--a-record] [--tags]`,
		Short: "Create floating IP address.",
		Long:  "Create floating IP address in speficied project.",
		Example: `  # Create a floating IP address in EU-Nord-1 location:
  cherryctl ip create -p <project_id> --region eu_nord_1`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			tagsArr := make(map[string]string)

			for _, kv := range tags {
				var k, v string
				tokens := strings.SplitN(kv, "=", 2)
				k = strings.TrimSpace(tokens[0])
				if len(tokens) != 1 {
					v = strings.TrimSpace(tokens[1])
				}

				tagsArr[k] = v
			}

			if targetHostname != "" {
				srvID, err := utils.ServerHostnameToID(targetHostname, projectID, c.ServerService)
				if err != nil {
					return errors.Wrap(err, "Could not find a target by hostname")
				}
				target = strconv.Itoa(srvID)
			} else if targetIP != "" {
				target = targetIP
			} else if targetID != 0 {
				target = strconv.Itoa(targetID)
			}

			request := &cherrygo.CreateIPAddress{
				Region:     region,
				PtrRecord:  ptrRecord,
				ARecord:    aRecord,
				TargetedTo: target,
				Tags:       &tagsArr,
			}

			i, _, err := c.Service.Create(projectID, request)
			if err != nil {
				return errors.Wrap(err, "Could not create IP address")
			}

			header := []string{"ID", "Address", "Cidr", "Type", "Region", "PTR record", "A record", "Tags"}
			data := make([][]string, 1)
			data[0] = []string{i.ID, i.Address, i.Cidr, i.Type, i.Region.Name, i.Region.Name, i.PtrRecord, fmt.Sprintf("%v", i.Tags)}

			return c.Out.Output(i, header, &data)
		},
	}

	ipCreateCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	ipCreateCmd.Flags().StringVarP(&region, "region", "", "", "Slug of the region from where IP address will requested.")
	ipCreateCmd.Flags().StringVarP(&targetHostname, "target-hostname", "", "", "The hostname of the server to assign the created IP to.")
	ipCreateCmd.Flags().IntVarP(&targetID, "target-id", "", 0, "The ID of the server to assign the created IP to.")
	ipCreateCmd.Flags().StringVarP(&targetIP, "target-ip-id", "", "", "Subnet or primary-ip type IP ID to route the created IP to.")
	ipCreateCmd.Flags().StringVarP(&ptrRecord, "ptr-record", "", "", "Slug of the region from where IP address will requested.")
	ipCreateCmd.Flags().StringVarP(&aRecord, "a-record", "", "", "Slug of the region from where IP address will requested.")
	ipCreateCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, `Tag or list of tags for the server: --tags="key=value,env=prod".`)

	ipCreateCmd.MarkFlagsMutuallyExclusive("target-hostname", "target-id", "target-ip-id")
	ipCreateCmd.MarkFlagsRequiredTogether("target-hostname", "project-id")

	ipCreateCmd.MarkFlagRequired("project-id")
	ipCreateCmd.MarkFlagRequired("region")

	return ipCreateCmd
}
