package backups

import (
	"strconv"
	"strings"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Methods() *cobra.Command {
	var backupID int
	servicesListCmd := &cobra.Command{
		Use:     `methods <backup_ID>`,
		Aliases: []string{"services", "protocols"},
		Args:    cobra.ExactArgs(1),
		Short:   "Retrieves backup storage access methods.",
		Long:    "Retrieves a list of available backup access methods.",
		Example: `  # Retrieves a list of backup storage methods:
		cherryctl backup methods 123`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if backID, err := strconv.Atoi(args[0]); err == nil {
				backupID = backID
			}

			storage, _, err := c.Service.Get(backupID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get a backup storage")
			}

			header := []string{"Name", "Host", "Username", "Password", "Port", "Processing", "Enabled", "Whitelist"}
			data := make([][]string, len(storage.Methods))

			for i, o := range storage.Methods {
				data[i] = []string{o.Name, o.Host, o.Username, o.Password, strconv.Itoa(o.Port), utils.BoolToYesNo(o.Processing), utils.BoolToYesNo(o.Enabled), strings.Join(o.WhiteList, ", ")}
			}

			return c.Out.Output(storage.Methods, header, &data)
		},
	}

	return servicesListCmd
}
