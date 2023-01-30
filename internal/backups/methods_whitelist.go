package backups

import (
	"strconv"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) MethodsWhitelist() *cobra.Command {
	var backupID int
	servicesWhitelistCmd := &cobra.Command{
		Use:     `methods-whitelist <backup_ID>`,
		Aliases: []string{"whitelist", "acl"},
		Args:    cobra.ExactArgs(1),
		Short:   "Retrieves a list of backup storage methods whitelist.",
		Long:    "Return information about whitelisted IP addresses and backup methods they are allowed to use.",
		Example: `  # Retrieves a list of backup storage methods whitelist:
		cherryctl backup methods-whitelist 123`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if backID, err := strconv.Atoi(args[0]); err == nil {
				backupID = backID
			}

			storage, _, err := c.Service.Get(backupID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get a backup storage")
			}

			header := []string{"IP Address", "FTP", "SMB", "NFS", "BORG"}
			data := make([][]string, len(storage.Rules))

			for i, o := range storage.Rules {
				data[i] = []string{o.IPAddress.Address, utils.BoolToYesNo(o.EnabledMethods.FTP), utils.BoolToYesNo(o.EnabledMethods.SMB), utils.BoolToYesNo(o.EnabledMethods.NFS), utils.BoolToYesNo(o.EnabledMethods.BORG)}
			}

			return c.Out.Output(storage.Rules, header, &data)
		},
	}

	return servicesWhitelistCmd
}
