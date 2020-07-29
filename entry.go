package main

import (
	"cherrygo"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {

	v1, err := readConfig(".cherry")
	// if err != nil {
	// 	log.Printf("Error reading config file, %s", err)
	// }

	// Init variables from configuration file
	var teamIDCfg int
	var projectIDCfg string

	if v1.IsSet("default_profile") {

		configDefaultProfile := v1.GetString("default_profile")

		teamIDCfg, projectIDCfg = getProfileConfig(configDefaultProfile)
	}

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
		Short: "Adds various objects",
		Long:  "Adds various objects",
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

	var cmdProfile = &cobra.Command{
		Use:   "profile",
		Short: "Manages profile settings from YAML config file $HOME/.cherry.yaml",
		Long:  "Manages profile settings from YAML config file $HOME/.cherry.yaml",
	}

	var cmdListPlans = &cobra.Command{
		Use:   "plans",
		Short: "List plans",
		Long:  "List plans for specified project",
		Run: func(cmd *cobra.Command, args []string) {

			if v1.IsSet("default_profile") {
				listPlans(c, teamIDCfg)
			} else {
				projectID, _ := cmd.Flags().GetInt("team-id")
				listPlans(c, projectID)
			}
		},
	}

	var cmdListProjects = &cobra.Command{
		Use:   "projects",
		Short: "List projects",
		Long:  "List projects for specified team",
		Run: func(cmd *cobra.Command, args []string) {

			if v1.IsSet("default_profile") {

				listProjects(c, teamIDCfg)
			} else {
				teamID, _ := cmd.Flags().GetInt("team-id")
				fmt.Printf("Read from flags: %v", projectIDCfg)
				listProjects(c, teamID)
			}
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

			if v1.IsSet("default_profile") {
				listServers(c, projectIDCfg)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				listServers(c, projectID)
			}

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

			if v1.IsSet("default_profile") {
				listIPAddresses(c, projectIDCfg)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				listIPAddresses(c, projectID)
			}

		},
	}

	var cmdListIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "List specific ip address",
		Long:  "List specific ip address",
		Run: func(cmd *cobra.Command, args []string) {

			ipID, _ := cmd.Flags().GetString("ip-id")

			if v1.IsSet("default_profile") {

				listIPAddress(c, projectIDCfg, ipID)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				listIPAddress(c, projectID, ipID)
			}

		},
	}

	var cmdAddIPAddress = &cobra.Command{
		Use:   "ip-address",
		Short: "Orders new floating IP address",
		Long:  "Orders new floating IP address",
		Run: func(cmd *cobra.Command, args []string) {

			aRecord, _ := cmd.Flags().GetString("a-record")
			ptrRecord, _ := cmd.Flags().GetString("ptr-record")
			routedTo, _ := cmd.Flags().GetString("routed-to")
			region, _ := cmd.Flags().GetString("region")

			if v1.IsSet("default_profile") {

				addIPAddress(c, projectIDCfg, aRecord, ptrRecord, routedTo, region)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				addIPAddress(c, projectID, aRecord, ptrRecord, routedTo, region)
			}

		},
	}

	var cmdAddServer = &cobra.Command{
		Use:   "server",
		Short: "Orders new server",
		Long:  "Orders new server",
		Run: func(cmd *cobra.Command, args []string) {

			hostname, _ := cmd.Flags().GetString("hostname")
			ipAddresses, _ := cmd.Flags().GetStringSlice("ip-addresses")
			sshKeys, _ := cmd.Flags().GetStringSlice("ssh-keys")
			image, _ := cmd.Flags().GetString("image")
			planID, _ := cmd.Flags().GetString("plan-id")
			region, _ := cmd.Flags().GetString("region")
			userData, _ := cmd.Flags().GetString("user-data")
			tags, _ := cmd.Flags().GetStringToString("tags")

			if v1.IsSet("default_profile") {
				addServer(c, projectIDCfg, hostname, ipAddresses, sshKeys, image, planID, region, userData, tags)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				addServer(c, projectID, hostname, ipAddresses, sshKeys, image, planID, region, userData, tags)
			}

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
			if v1.IsSet("default_profile") {

				createProject(c, teamIDCfg, projectName)
			} else {
				teamID, _ := cmd.Flags().GetInt("team-id")
				createProject(c, teamID, projectName)
			}
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

			ipID, _ := cmd.Flags().GetString("ip-id")
			if v1.IsSet("default_profile") {

				deleteIPAddress(c, projectIDCfg, ipID)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				deleteIPAddress(c, projectID, ipID)
			}
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

			if v1.IsSet("default_profile") {

				updateIPAddress(c, ptrRecord, aRecord,
					routedTo, routedToHostname, routedToServerIP, routedToServerID,
					assignedTo, floatingID, floatingIP, projectIDCfg)
			} else {
				projectID, _ := cmd.Flags().GetString("project-id")
				updateIPAddress(c, ptrRecord, aRecord,
					routedTo, routedToHostname, routedToServerIP, routedToServerID,
					assignedTo, floatingID, floatingIP, projectID)
			}

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

	var cmdUpdateServer = &cobra.Command{
		Use:   "server",
		Short: "Updates specified server",
		Long:  "Updates specified server",
		Run: func(cmd *cobra.Command, args []string) {

			serverID, _ := cmd.Flags().GetString("server-id")
			tags, _ := cmd.Flags().GetStringToString("tags")
			updateServer(c, tags, serverID)
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

	var cmdProfileSetDefault = &cobra.Command{
		Use:   "set",
		Short: "set default profile",
		Long:  "set default profile",
		Run: func(cmd *cobra.Command, args []string) {

			defaultProfile, _ := cmd.Flags().GetString("default-profile")
			setDefaultProfile(defaultProfile)
		},
	}

	var rootCmd = &cobra.Command{Use: "cherry-cloud-cli"}
	rootCmd.AddCommand(cmdList, cmdAdd, cmdRemove, cmdUpdate, cmdPower, cmdProfile)

	cmdList.AddCommand(cmdListTeams, cmdListPlans, cmdListProjects, cmdListImages, cmdListSSHKeys, cmdListSSHKey, cmdListServers, cmdListServer, cmdListIPAddresses, cmdListIPAddress, cmdListProject)

	cmdListPlans.Flags().IntP("team-id", "t", 0, "Provide team-id")
	if !v1.IsSet("default_profile") {
		cmdListProjects.MarkFlagRequired("team-id")
	}

	cmdListImages.Flags().IntP("plan-id", "p", 0, "Provide plan-id")
	cmdListImages.MarkFlagRequired("plan-id")

	cmdListProjects.Flags().IntP("team-id", "t", 0, "Provide team-id")
	if !v1.IsSet("default_profile") {
		cmdListProjects.MarkFlagRequired("team-id")
	}

	cmdListProject.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListProject.MarkFlagRequired("project-id")

	cmdListServers.Flags().StringP("project-id", "p", "", "Provide project-id")
	if !v1.IsSet("default_profile") {
		cmdListServers.MarkFlagRequired("project-id")
	}

	cmdListServer.Flags().StringP("server-id", "s", "", "Provide server-id")
	cmdListServer.MarkFlagRequired("server-id")

	cmdListSSHKey.Flags().StringP("key-id", "k", "", "Provide key-id")
	cmdListSSHKey.MarkFlagRequired("key-id")

	cmdListIPAddresses.Flags().StringP("project-id", "p", "", "Provide project-id")
	if !v1.IsSet("default_profile") {
		cmdListIPAddresses.MarkFlagRequired("project-id")
	}

	cmdListIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdListIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip-id")
	if !v1.IsSet("default_profile") {
		cmdListIPAddress.MarkFlagRequired("project-id")
	}

	cmdListIPAddress.MarkFlagRequired("ip-id")

	// Add section
	cmdAdd.AddCommand(cmdAddIPAddress, cmdAddSSHKey, cmdAddServer, cmdAddProject)

	// Add new ip address section
	cmdAddIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdAddIPAddress.Flags().StringP("a-record", "a", "", "Provide a-record")
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
	cmdAddServer.Flags().StringP("user-data", "u", "", "Provide blob of user-data in base64")
	var tags = make(map[string]string)
	cmdAddServer.Flags().StringToStringVar(&tags, "tags", nil, "Provide key/value for tags: env=prod,name=node1")
	var keySlice, ipSlice []string
	cmdAddServer.Flags().StringSliceP("ssh-keys", "k", keySlice, "Provide ssh-keys")
	cmdAddServer.Flags().StringSliceP("ip-addresses", "d", ipSlice, "Provide ip-addresses")

	// Required flags for creating a new server
	if !v1.IsSet("default_profile") {
		cmdAddServer.MarkFlagRequired("project-id")
	}
	cmdAddServer.MarkFlagRequired("region")
	cmdAddServer.MarkFlagRequired("image")
	cmdAddServer.MarkFlagRequired("plan-id")
	cmdAddServer.MarkFlagRequired("hostname")

	cmdAddProject.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdAddProject.Flags().StringP("project-name", "p", "", "Provide project-name")

	// Required flags for creating new project
	if !v1.IsSet("default_profile") {
		cmdAddProject.MarkFlagRequired("team-id")
	}
	cmdAddProject.MarkFlagRequired("project-name")

	// Remove section
	cmdRemove.AddCommand(cmdRemoveSSHKey, cmdRemoveIPAddress, cmdRemoveServer, cmdRemoveProject)

	cmdRemoveSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for removal")
	cmdRemoveSSHKey.MarkFlagRequired("key-id")

	cmdRemoveIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip id for removal")
	cmdRemoveIPAddress.Flags().IntP("project-id", "p", 0, "Provide project-id")
	cmdRemoveIPAddress.MarkFlagRequired("ip-id")
	if !v1.IsSet("default_profile") {
		cmdRemoveIPAddress.MarkFlagRequired("project-id")
	}

	cmdRemoveServer.Flags().StringP("server-id", "s", "", "Provide server id for removal")
	cmdRemoveServer.MarkFlagRequired("server-id")

	cmdRemoveProject.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdRemoveProject.MarkFlagRequired("project-id")

	// Update section
	cmdUpdate.AddCommand(cmdUpdateSSHKey, cmdUpdateIPAddress, cmdUpdateProject, cmdUpdateServer)

	cmdUpdateSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for update")
	cmdUpdateSSHKey.Flags().StringP("key-label", "l", "", "Provide new label for key")

	// Update IP address section
	cmdUpdateIPAddress.Flags().StringP("floating-id", "i", "", "Provide floating ip id for update")
	cmdUpdateIPAddress.Flags().StringP("floating-ip", "f", "", "Provide floating ip for update")
	cmdUpdateIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdUpdateIPAddress.Flags().StringP("a-record", "a", "", "Provide a-record")
	cmdUpdateIPAddress.Flags().StringP("ptr-record", "r", "", "Provide ptr-record")
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

	// Update Server section
	cmdUpdateServer.Flags().StringP("server-id", "s", "", "Provide server-id")
	cmdUpdateServer.Flags().StringToStringVar(&tags, "tags", nil, "Provide key/value for tags: env=prod,name=node1")

	// Required flags for updating server
	cmdUpdateServer.MarkFlagRequired("server-id")
	cmdUpdateServer.MarkFlagRequired("tags")

	// Power section
	cmdPower.AddCommand(cmdPowerOnServer, cmdPowerOffServer, cmdRebootServer)
	cmdPowerOnServer.Flags().StringP("server-id", "s", "", "Provide server id for power on")
	cmdPowerOnServer.MarkFlagRequired("server-id")

	cmdPowerOffServer.Flags().StringP("server-id", "s", "", "Provide server id for power on")
	cmdPowerOffServer.MarkFlagRequired("server-id")

	cmdRebootServer.Flags().StringP("server-id", "s", "", "Provide server id for reboot")
	cmdRebootServer.MarkFlagRequired("server-id")

	// Profile section
	cmdProfile.AddCommand(cmdProfileSetDefault)
	cmdProfileSetDefault.Flags().StringP("default-profile", "p", "", "Provide default profile name to use")

	// Required flags for profile section
	cmdProfileSetDefault.MarkFlagRequired("default-profile")

	rootCmd.Execute()

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
