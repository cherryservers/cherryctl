package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func main() {

	// Let's create new client
	c, err := cherrygo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Lists various objects",
		Long:  "Lists various objects (teams, projects, images, servers)",
	}

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Ads various objects",
		Long:  "Ads various objects",
	}

	var cmdRemove = &cobra.Command{
		Use:   "remove",
		Short: "Removes various objects",
		Long:  "Removes various objects",
	}

	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "Updates various objects",
		Long:  "Updates various objects",
	}

	var cmdPower = &cobra.Command{
		Use:   "power",
		Short: "Manages power on servers",
		Long:  "Manages power on servers",
	}

	var cmdListPlans = &cobra.Command{
		Use:   "plans",
		Short: "List plans",
		Long:  "List plans for specified project",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetInt("team-id")
			listPlans(c, projectID)
		},
	}

	var cmdListProjects = &cobra.Command{
		Use:   "projects",
		Short: "List projects",
		Long:  "List projects for specified team",
		Run: func(cmd *cobra.Command, args []string) {
			teamID, _ := cmd.Flags().GetInt("team-id")
			listProjects(c, teamID)
		},
	}

	var cmdListProject = &cobra.Command{
		Use:   "project",
		Short: "List project",
		Long:  "List project",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetString("project-id")
			listProject(c, projectID)
		},
	}

	var cmdListImages = &cobra.Command{
		Use:   "images",
		Short: "List images",
		Long:  "List images for specified plan",
		Run: func(cmd *cobra.Command, args []string) {
			planID, _ := cmd.Flags().GetInt("plan-id")
			listImages(c, planID)
		},
	}

	var cmdListTeams = &cobra.Command{
		Use:   "teams",
		Short: "List teams",
		Long:  "List teams",
		Run: func(cmd *cobra.Command, args []string) {
			listTeams(c)
		},
	}

	var cmdListSSHKeys = &cobra.Command{
		Use:   "ssh-keys",
		Short: "List ssh keys",
		Long:  "List ssh keys",
		Run: func(cmd *cobra.Command, args []string) {
			listSSHkeys(c)
		},
	}

	var cmdListSSHKey = &cobra.Command{
		Use:   "ssh-key",
		Short: "List ssh key",
		Long:  "List ssh key",
		Run: func(cmd *cobra.Command, args []string) {
			keyID, _ := cmd.Flags().GetString("key-id")
			listSSHkey(c, keyID)
		},
	}

	var cmdListServers = &cobra.Command{
		Use:   "servers",
		Short: "List servers",
		Long:  "List servers for specified project",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetString("project-id")
			listServers(c, projectID)
		},
	}

	var cmdListServer = &cobra.Command{
		Use:   "server",
		Short: "List specific server",
		Long:  "List specific server",
		Run: func(cmd *cobra.Command, args []string) {
			serverID, _ := cmd.Flags().GetString("server-id")
			listServer(c, serverID)
		},
	}

	var cmdListIPAddresses = &cobra.Command{
		Use:   "ip-addresses",
		Short: "List ip addresses",
		Long:  "List ip addresses",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetString("project-id")
			listIPAddresses(c, projectID)
		},
	}

	var cmdListIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "List specific ip address",
		Long:  "List specific ip address",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetString("project-id")
			ipID, _ := cmd.Flags().GetString("ip-id")
			listIPAddress(c, projectID, ipID)
		},
	}

	var cmdAddIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "Orders new floating IP address",
		Long:  "Orders new floating IP address",
		Run: func(cmd *cobra.Command, args []string) {
			projectID, _ := cmd.Flags().GetString("project-id")
			aRecord, _ := cmd.Flags().GetString("a-record")
			ptrRecord, _ := cmd.Flags().GetString("ptr-record")
			routedTo, _ := cmd.Flags().GetString("routed-to")
			region, _ := cmd.Flags().GetString("region")
			addIPAddress(c, projectID, aRecord, ptrRecord, routedTo, region)
		},
	}

	var cmdAddServer = &cobra.Command{
		Use:   "server",
		Short: "Orders new server",
		Long:  "Orders new server",
		Run: func(cmd *cobra.Command, args []string) {

			projectID, _ := cmd.Flags().GetString("project-id")
			hostname, _ := cmd.Flags().GetString("hostname")
			ipAddresses, _ := cmd.Flags().GetStringSlice("ip-addresses")
			sshKeys, _ := cmd.Flags().GetStringSlice("ssh-keys")
			image, _ := cmd.Flags().GetString("image")
			planID, _ := cmd.Flags().GetString("plan-id")
			region, _ := cmd.Flags().GetString("region")
			addServer(c, projectID, hostname, ipAddresses, sshKeys, image, planID, region)
		},
	}

	var cmdAddSSHKey = &cobra.Command{
		Use:   "ssh-key",
		Short: "Adds new ssh key",
		Long:  "Adds new ssh key",
		Run: func(cmd *cobra.Command, args []string) {

			label, _ := cmd.Flags().GetString("key-label")
			key, _ := cmd.Flags().GetString("raw-key")
			keyPath, _ := cmd.Flags().GetString("key-path")
			createSSHKey(c, label, key, keyPath)
		},
	}

	var cmdAddProject = &cobra.Command{
		Use:   "project",
		Short: "Adds new project",
		Long:  "Adds new project",
		Run: func(cmd *cobra.Command, args []string) {

			projectName, _ := cmd.Flags().GetString("project-name")
			teamID, _ := cmd.Flags().GetInt("team-id")

			createProject(c, teamID, projectName)
		},
	}

	var cmdRemoveSSHKey = &cobra.Command{
		Use:   "ssh-key",
		Short: "Removes specified ssh key",
		Long:  "Removes specified ssh key",
		Run: func(cmd *cobra.Command, args []string) {

			keyID, _ := cmd.Flags().GetString("key-id")
			deleteSSHKey(c, keyID)
		},
	}

	var cmdRemoveServer = &cobra.Command{
		Use:   "server",
		Short: "Removes specified server",
		Long:  "Removes specified server",
		Run: func(cmd *cobra.Command, args []string) {

			serverID, _ := cmd.Flags().GetString("server-id")
			deleteServer(c, serverID)
		},
	}

	var cmdRemoveIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "Removes specified ip address",
		Long:  "Removes specified ip address",
		Run: func(cmd *cobra.Command, args []string) {

			projectID, _ := cmd.Flags().GetString("project-id")
			ipID, _ := cmd.Flags().GetString("ip-id")
			deleteIPAddress(c, projectID, ipID)
		},
	}

	var cmdRemoveProject = &cobra.Command{
		Use:   "project",
		Short: "Removes specified project",
		Long:  "Removes specified project",
		Run: func(cmd *cobra.Command, args []string) {

			projectID, _ := cmd.Flags().GetString("project-id")
			deleteProject(c, projectID)
		},
	}

	var cmdUpdateSSHKey = &cobra.Command{
		Use:   "ssh-key",
		Short: "Updates specified ssh key",
		Long:  "Updates specified ssh key",
		Run: func(cmd *cobra.Command, args []string) {

			keyID, _ := cmd.Flags().GetString("key-id")
			keyLabel, _ := cmd.Flags().GetString("key-label")
			updateSSHKey(c, keyID, keyLabel)
		},
	}

	var cmdUpdateIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "Updates specified ip address",
		Long:  "Updates specified ip address",
		Run: func(cmd *cobra.Command, args []string) {

			floatingID, _ := cmd.Flags().GetString("floating-id")
			floatingIP, _ := cmd.Flags().GetString("floating-ip")
			projectID, _ := cmd.Flags().GetString("project-id")
			aRecord, _ := cmd.Flags().GetString("a-record")
			ptrRecord, _ := cmd.Flags().GetString("ptr-record")
			routedTo, _ := cmd.Flags().GetString("routed-to")
			assignedTo, _ := cmd.Flags().GetString("assigned-to")
			routedToHostname, _ := cmd.Flags().GetString("routed-to-hostname")
			routedToServerIP, _ := cmd.Flags().GetString("routed-to-server-ip")
			routedToServerID, _ := cmd.Flags().GetString("routed-to-server-id")

			err := validateArgs(cmd, "cmdUpdateIPAddress")
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			updateIPAddress(c, ptrRecord, aRecord,
				routedTo, routedToHostname, routedToServerIP, routedToServerID,
				assignedTo, floatingID, floatingIP, projectID)
		},
	}

	var cmdUpdateProject = &cobra.Command{
		Use:   "project",
		Short: "Updates specified project",
		Long:  "Updates specified project",
		Run: func(cmd *cobra.Command, args []string) {

			projectID, _ := cmd.Flags().GetString("project-id")
			projectName, _ := cmd.Flags().GetString("project-name")
			updateProject(c, projectID, projectName)
		},
	}

	var cmdPowerOnServer = &cobra.Command{
		Use:   "on",
		Short: "Powers server ON",
		Long:  "Powers server ON",
		Run: func(cmd *cobra.Command, args []string) {

			serverID, _ := cmd.Flags().GetString("server-id")
			powerOn(c, serverID)
		},
	}

	var cmdPowerOffServer = &cobra.Command{
		Use:   "off",
		Short: "Powers server OFF",
		Long:  "Powers server OFF",
		Run: func(cmd *cobra.Command, args []string) {

			serverID, _ := cmd.Flags().GetString("server-id")
			powerOff(c, serverID)
		},
	}

	var cmdRebootServer = &cobra.Command{
		Use:   "reboot",
		Short: "Reboots the server",
		Long:  "Reboots the server",
		Run: func(cmd *cobra.Command, args []string) {

			serverID, _ := cmd.Flags().GetString("server-id")
			rebootServer(c, serverID)
		},
	}

	var rootCmd = &cobra.Command{Use: "cherry-cloud-cli"}
	rootCmd.AddCommand(cmdList, cmdAdd, cmdRemove, cmdUpdate, cmdPower)

	cmdList.AddCommand(cmdListTeams, cmdListPlans, cmdListProjects, cmdListImages, cmdListSSHKeys, cmdListSSHKey, cmdListServers, cmdListServer, cmdListIPAddresses, cmdListIPAddress, cmdListProject)

	cmdListPlans.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdListPlans.MarkFlagRequired("team-id")

	cmdListImages.Flags().IntP("plan-id", "p", 0, "Provide plan-id")
	cmdListImages.MarkFlagRequired("plan-id")

	cmdListProjects.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdListProjects.MarkFlagRequired("team-id")

	cmdListProject.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListProject.MarkFlagRequired("project-id")

	cmdListServers.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListServers.MarkFlagRequired("project-id")

	cmdListServer.Flags().StringP("server-id", "s", "", "Provide server-id")
	cmdListServer.MarkFlagRequired("server-id")

	cmdListSSHKey.Flags().StringP("key-id", "k", "", "Provide key-id")
	cmdListSSHKey.MarkFlagRequired("key-id")

	cmdListIPAddresses.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListIPAddresses.MarkFlagRequired("project-id")

	cmdListIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip-id")
	cmdListIPAddress.MarkFlagRequired("project-id")
	cmdListIPAddress.MarkFlagRequired("ip-id")

	// Add section
	cmdAdd.AddCommand(cmdAddIPAddress, cmdAddSSHKey, cmdAddServer, cmdAddProject)

	// Add new ip address section
	cmdAddIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdAddIPAddress.Flags().StringP("a-record", "a", "a-record.example.com", "Provide a-record")
	cmdAddIPAddress.Flags().StringP("ptr-record", "r", "ptr-record.examples.com", "Provide ptr-record")
	cmdAddIPAddress.Flags().StringP("region", "g", "EU-East-1", "Provide region")
	cmdAddIPAddress.Flags().StringP("routed-to", "t", "", "Provide ipaddress_id to route to")

	// Add ssh key section
	cmdAddSSHKey.Flags().StringP("key-label", "l", "ssh-key-label", "Provide ssh key label")
	cmdAddSSHKey.Flags().StringP("raw-key", "k", "", "Provide ssh raw key")
	cmdAddSSHKey.Flags().StringP("key-path", "f", "", "Provide path to ssh key")

	// Add new server section
	cmdAddServer.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdAddServer.Flags().StringP("hostname", "s", "server-name.examples.com", "Provide hostname")
	cmdAddServer.Flags().StringP("region", "g", "EU-East-1", "Provide region")
	cmdAddServer.Flags().StringP("image", "i", "", "Provide image")
	cmdAddServer.Flags().StringP("plan-id", "l", "", "Provide plan-id")
	var aa []string
	var zz []string
	cmdAddServer.Flags().StringSliceP("ssh-keys", "k", aa, "Provide ssh-keys")
	cmdAddServer.Flags().StringSliceP("ip-addresses", "d", zz, "Provide image")

	// Required flags for creating a new server
	cmdAddServer.MarkFlagRequired("project-id")
	cmdAddServer.MarkFlagRequired("region")
	cmdAddServer.MarkFlagRequired("image")
	cmdAddServer.MarkFlagRequired("plan-id")
	cmdAddServer.MarkFlagRequired("hostname")

	cmdAddProject.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdAddProject.Flags().StringP("project-name", "p", "", "Provide project-name")

	// Required flags for creating new project
	cmdAddProject.MarkFlagRequired("team-id")
	cmdAddProject.MarkFlagRequired("project-name")

	// Remove section
	cmdRemove.AddCommand(cmdRemoveSSHKey, cmdRemoveIPAddress, cmdRemoveServer, cmdRemoveProject)

	cmdRemoveSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for removal")
	cmdRemoveSSHKey.MarkFlagRequired("key-id")

	cmdRemoveIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip id for removal")
	cmdRemoveIPAddress.Flags().IntP("project-id", "p", 0, "Provide project-id")
	cmdRemoveIPAddress.MarkFlagRequired("ip-id")
	cmdRemoveIPAddress.MarkFlagRequired("project-id")

	cmdRemoveServer.Flags().StringP("server-id", "s", "", "Provide server id for removal")
	cmdRemoveServer.MarkFlagRequired("server-id")

	cmdRemoveProject.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdRemoveProject.MarkFlagRequired("project-id")

	// Update section
	cmdUpdate.AddCommand(cmdUpdateSSHKey, cmdUpdateIPAddress, cmdUpdateProject)

	cmdUpdateSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for update")
	cmdUpdateSSHKey.Flags().StringP("key-label", "l", "", "Provide new label for key")

	// Update IP address section
	cmdUpdateIPAddress.Flags().StringP("floating-id", "i", "", "Provide floating ip id for update")
	cmdUpdateIPAddress.Flags().StringP("floating-ip", "f", "", "Provide floating ip for update")
	cmdUpdateIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdUpdateIPAddress.Flags().StringP("a-record", "a", "a-record.example.com", "Provide a-record")
	cmdUpdateIPAddress.Flags().StringP("ptr-record", "r", "ptr-record.examples.com", "Provide ptr-record")
	cmdUpdateIPAddress.Flags().StringP("routed-to", "t", "", "Provide ipaddress_id to route to")
	cmdUpdateIPAddress.Flags().StringP("routed-to-hostname", "n", "", "Provide hostname of the server to route to")
	cmdUpdateIPAddress.Flags().StringP("routed-to-server-ip", "s", "", "Provide primary ip of the server to route to")
	cmdUpdateIPAddress.Flags().StringP("routed-to-server-id", "d", "", "Provide id of the server to route to")

	// Required flags for updating IP address
	cmdUpdateIPAddress.MarkFlagRequired("ip-id")
	cmdUpdateIPAddress.MarkFlagRequired("project-id")

	// Update Project section
	cmdUpdateProject.Flags().StringP("project-id", "i", "", "Provide project-id")
	cmdUpdateProject.Flags().StringP("project-name", "p", "", "Provide project-name")

	// Required flags for updating project
	cmdUpdateProject.MarkFlagRequired("project-id")
	cmdUpdateProject.MarkFlagRequired("project-name")

	// Power section
	cmdPower.AddCommand(cmdPowerOnServer, cmdPowerOffServer, cmdRebootServer)
	cmdPowerOnServer.Flags().StringP("server-id", "s", "", "Provide server id for power on")
	cmdPowerOnServer.MarkFlagRequired("server-id")

	cmdPowerOffServer.Flags().StringP("server-id", "s", "", "Provide server id for power on")
	cmdPowerOffServer.MarkFlagRequired("server-id")

	cmdRebootServer.Flags().StringP("server-id", "s", "", "Provide server id for reboot")
	cmdRebootServer.MarkFlagRequired("server-id")

	rootCmd.Execute()

}

func listPlans(c *cherrygo.Client, projectID int) {

	plans, _, err := c.Plans.List(projectID)
	if err != nil {
		log.Fatalf("Plans error: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t-------------\t----------\t-------------\t---\t-------\t---\n")
	fmt.Fprintf(tw, "Plan ID\tPlan name\tPlan price\tCPU\tRAM\tRegions\tQty\n")
	fmt.Fprintf(tw, "-------\t-------------\t----------\t-------------\t---\t-------\t---\n")
	for _, p := range plans {

		for _, r := range p.AvailableRegions {

			var planPrice float32
			for _, price := range p.Pricing {
				if price.Unit == "Hourly" {
					planPrice = price.Price
					// Prints most expensive plans
					if planPrice > 0 {
						fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
							p.ID, p.Name, planPrice, p.Specs.Cpus.Name, p.Specs.Memory.Total, r.Name, r.StockQty)

					}
				}
			}

		}

	}
	fmt.Fprintf(tw, "-------\t-------------\t----------\t-------------\t---\t-------\t---\n")
	tw.Flush()
}

func listImages(c *cherrygo.Client, planID int) {

	images, _, err := c.Images.List(planID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n--------\t----------\t-----------\n")
	fmt.Fprintf(tw, "Image ID\tImage name\tImage price\n")
	fmt.Fprintf(tw, "--------\t----------\t-----------\n")

	for _, i := range images {
		fmt.Fprintf(tw, "%v\t%v\t%v\n",
			i.ID, i.Name, i.Pricing.Price)
	}
	fmt.Fprintf(tw, "--------\t----------\t-----------\n")
	tw.Flush()
}

func listTeams(c *cherrygo.Client) {

	// Let's get and print info about Teams
	teams, _, err := c.Teams.List()
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t---------\t---------------\t------------\t-------\n")
	fmt.Fprintf(tw, "Team ID\tTeam name\tPromo remaining\tPromo usage\tPricing\n")
	fmt.Fprintf(tw, "-------\t---------\t---------------\t------------\t-------\n")

	for _, t := range teams {

		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n",
			t.ID, t.Name, t.Credit.Promo.Remaining, t.Credit.Promo.Usage, t.Credit.Resources.Pricing.Price)
	}
	fmt.Fprintf(tw, "-------\t---------\t---------------\t------------\t-------\n")
	tw.Flush()
}

func printDebugInfo(c *cherrygo.Client) {

	log.Println(c.AuthToken)
	log.Println(c.BaseURL)
	log.Println(c.UserAgent)

}

func validateArgs(cmd *cobra.Command, command string) error {

	switch {
	case command == "cmdUpdateIPAddress":

		routedTo, _ := cmd.Flags().GetString("routed-to")
		routedToHostname, _ := cmd.Flags().GetString("routed-to-hostname")
		routedToServerIP, _ := cmd.Flags().GetString("routed-to-server-ip")
		routedToServerID, _ := cmd.Flags().GetString("routed-to-server-id")

		switch {
		case len(routedTo) > 0:
			if len(routedToHostname) > 0 {
				return fmt.Errorf("--routed-to and --routed-to-hostname are mutually exclusive")
			}
			if len(routedToServerID) > 0 {
				return fmt.Errorf("--routed-to and --routed-to-server-id are mutually exclusive")
			}
			if len(routedToServerIP) > 0 {
				return fmt.Errorf("--routed-to and --routed-to-server-ip are mutually exclusive")
			}
		case len(routedToHostname) > 0:
			if len(routedToServerID) > 0 {
				return fmt.Errorf("--routed-to-hostname and --routed-to-server-id are mutually exclusive")
			}
			if len(routedToServerIP) > 0 {
				return fmt.Errorf("--routed-to-hostname and --routed-to-server-ip are mutually exclusive")
			}
		case len(routedToServerIP) > 0:
			if len(routedToServerID) > 0 {
				return fmt.Errorf("--routed-to-server-ip and --routed-to-server-id are mutually exclusive")
			}
		}

	}

	return nil
}
