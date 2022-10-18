package backups

import (
	"strconv"
	"strings"

	"github.com/cherryservers/cherryctl/internal/utils"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) UpdateMethod() *cobra.Command {
	var (
		backupID    int
		serviceName string
		ipWhitelist []string
		enable      bool
		disable     bool
	)
	backupUpdateCmd := &cobra.Command{
		Use:     `update-method <backup_ID> --method-name <string> [--enable] [--disable] [--whitelist <ip_addresses>]`,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"update-method"},
		Short:   "Update a backup storage access method.",
		Long:    "Enable or disable the selected backup access method or set a list of available IP addresses allowed to use this method.",
		Example: `  # Enable FTP protocol for your backup storage:
  cherryctl backup update-method 12345 --method-name FTP --enable`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if backID, err := strconv.Atoi(args[0]); err == nil {
				backupID = backID
			}

			request := &cherrygo.UpdateBackupMethod{
				BackupStorageID:  backupID,
				BackupMethodName: serviceName,
			}

			if len(ipWhitelist) == 0 || (len(ipWhitelist) > 0 && ipWhitelist[0] != "0") {
				request.Whitelist = ipWhitelist
			}

			if enable {
				request.Enabled = true
			} else if disable {
				request.Enabled = false
			}

			services, _, err := c.Service.UpdateBackupMethod(request)
			if err != nil {
				return errors.Wrap(err, "Could not update backup storage access method")
			}

			header := []string{"Name", "Host", "Username", "Password", "Port", "Processing", "Enabled", "Whitelist"}
			data := make([][]string, len(services))

			for i, o := range services {
				data[i] = []string{o.Name, o.Host, o.Username, o.Password, strconv.Itoa(o.Port), utils.BoolToYesNo(o.Processing), utils.BoolToYesNo(o.Enabled), strings.Join(o.WhiteList, ", ")}
			}
			return c.Out.Output(services, header, &data)
		},
	}

	backupUpdateCmd.Flags().StringVarP(&serviceName, "method-name", "n", "", "Backup access method name.")
	backupUpdateCmd.Flags().StringSliceVarP(&ipWhitelist, "whitelist", "", []string{"0"}, "A comma separated list of IP addresses to be whitelisted for access via this backup method.")
	backupUpdateCmd.Flags().BoolVarP(&enable, "enable", "e", false, "Enable method.")
	backupUpdateCmd.Flags().BoolVarP(&disable, "disable", "d", false, "Disable method.")

	backupUpdateCmd.MarkFlagRequired("method-name")

	return backupUpdateCmd
}
