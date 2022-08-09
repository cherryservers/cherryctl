package storages

import (
	"fmt"
	"strconv"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		projectID   int
		description string
		size        int
		region      string
	)
	storageCreateCmd := &cobra.Command{
		Use:   `create [-p <project_id>] --size <gigabytes> --region <region_slug> [--description]`,
		Short: "Create storage.",
		Long:  "Create storage in speficied project.",
		Example: `  # Create storage volume with 500GB space in EU-Nord-1 location:
  cherryctl storage create -p 12345 --size 500 --region eu_nord_1`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := &cherrygo.CreateStorage{
				ProjectID:   projectID,
				Description: description,
				Size:        size,
				Region:      region,
			}

			o, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not create storage")
			}

			header := []string{"ID", "Size", "Region", "Description"}
			data := make([][]string, 1)
			data[0] = []string{strconv.Itoa(o.ID), fmt.Sprintf("%d %s", o.Size, o.Unit), o.Region.Name, o.Description}

			return c.Out.Output(o, header, &data)
		},
	}

	storageCreateCmd.Flags().IntVarP(&projectID, "project-id", "p", 0, "The project's ID.")
	storageCreateCmd.Flags().IntVarP(&size, "size", "", 0, "Storage volume size in gigabytes.")
	storageCreateCmd.Flags().StringVarP(&region, "region", "", "", "Slug of the region.")
	storageCreateCmd.Flags().StringVarP(&description, "description", "", "", "Storage description.")

	storageCreateCmd.MarkFlagRequired("project-id")
	storageCreateCmd.MarkFlagRequired("size")
	storageCreateCmd.MarkFlagRequired("region")

	return storageCreateCmd
}
