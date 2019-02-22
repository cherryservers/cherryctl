package main

import (
	"cherrygo"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

			ipID, _ := cmd.Flags().GetString("ip-id")
			projectID, _ := cmd.Flags().GetString("project-id")
			aRecord, _ := cmd.Flags().GetString("a-record")
			ptrRecord, _ := cmd.Flags().GetString("ptr-record")
			routedTo, _ := cmd.Flags().GetString("routed-to")
			assignedTo, _ := cmd.Flags().GetString("assigned-to")
			updateIPAddress(c, ptrRecord, aRecord, routedTo, assignedTo, ipID, projectID)
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

	cmdList.AddCommand(cmdListTeams, cmdListPlans, cmdListProjects, cmdListImages, cmdListSSHKeys, cmdListSSHKey, cmdListServers, cmdListServer, cmdListIPAddresses, cmdListIPAddress)

	cmdListPlans.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdListPlans.MarkFlagRequired("team-id")

	cmdListImages.Flags().IntP("plan-id", "p", 0, "Provide plan-id")
	cmdListImages.MarkFlagRequired("plan-id")

	cmdListProjects.Flags().IntP("team-id", "t", 0, "Provide team-id")
	cmdListProjects.MarkFlagRequired("team-id")

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
	cmdAdd.AddCommand(cmdAddIPAddress, cmdAddSSHKey, cmdAddServer)

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

	// Remove section
	cmdRemove.AddCommand(cmdRemoveSSHKey, cmdRemoveIPAddress, cmdRemoveServer)

	cmdRemoveSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for removal")
	cmdRemoveSSHKey.MarkFlagRequired("key-id")

	cmdRemoveIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip id for removal")
	cmdRemoveIPAddress.Flags().IntP("project-id", "p", 0, "Provide project-id")
	cmdRemoveIPAddress.MarkFlagRequired("ip-id")
	cmdRemoveIPAddress.MarkFlagRequired("project-id")

	cmdRemoveServer.Flags().StringP("server-id", "s", "", "Provide server id for removal")
	cmdRemoveServer.MarkFlagRequired("server-id")

	// Update section
	cmdUpdate.AddCommand(cmdUpdateSSHKey, cmdUpdateIPAddress)

	cmdUpdateSSHKey.Flags().StringP("key-id", "k", "", "Provide ssh key id for update")
	cmdUpdateSSHKey.Flags().StringP("key-label", "l", "", "Provide new label for key")

	// Update IP address section
	cmdUpdateIPAddress.Flags().StringP("ip-id", "i", "", "Provide ip id for update")
	cmdUpdateIPAddress.Flags().StringP("project-id", "p", "", "Provide project-id")
	cmdUpdateIPAddress.Flags().StringP("a-record", "a", "a-record.example.com", "Provide a-record")
	cmdUpdateIPAddress.Flags().StringP("ptr-record", "r", "ptr-record.examples.com", "Provide ptr-record")
	cmdUpdateIPAddress.Flags().StringP("routed-to", "t", "", "Provide ipaddress_id to route to")

	// Required flags for updating IP address
	cmdUpdateIPAddress.MarkFlagRequired("ip-id")
	cmdUpdateIPAddress.MarkFlagRequired("project-id")

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

func addIPAddress(c *cherrygo.Client, projectID, aRecord, ptrRecord, routedTo, region string) {

	addIPRequest := cherrygo.CreateIPAddress{
		ARecord:   aRecord,
		PtrRecord: ptrRecord,
		Region:    region,
		RoutedTo:  routedTo,
	}

	ip, _, err := c.IPAddress.Create(projectID, &addIPRequest)
	if err != nil {
		log.Fatal("Error while ordering new floating IP address: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "IP ID\tAddress\tCidr\tPTR\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n", ip.ID, ip.Address, ip.Cidr, ip.PtrRecord, ip.Type)
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	tw.Flush()
}

func addServer(c *cherrygo.Client, projectID string, hostname string, ipaddresses []string, sshKeys []string, image string, planID string, region string) {

	addServerRequest := cherrygo.CreateServer{
		ProjectID:   projectID,
		Hostname:    hostname,
		Image:       image,
		Region:      region,
		SSHKeys:     sshKeys,
		IPAddresses: ipaddresses,
		PlanID:      planID,
	}

	server, _, err := c.Server.Create(projectID, &addServerRequest)
	if err != nil {
		log.Fatal("Error while creating new server: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "Server ID\tHostname\tName\tPTR\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n", server.ID, server.Hostname, server.Name, server.Image, server.State)
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	tw.Flush()
}

func listServers(c *cherrygo.Client, projectID string) {

	// Needs project id to be passed
	servers, _, err := c.Servers.List(projectID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t---\t----------\t--------\t------\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tCPU\tIP address\tSSH keys\tRegion\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----------\t--------\t------\t-----\n")

	for _, srv := range servers {

		for _, i := range srv.IPAddresses {
			if len(srv.SSHKeys) == 0 {
				fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
					srv.ID, srv.Name, srv.Hostname, srv.Image, srv.Plans.Specs.Cpus.Frequency, i.Address, nil, srv.Region.Name, srv.State)
			} else {
				for _, k := range srv.SSHKeys {
					fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
						srv.ID, srv.Name, srv.Hostname, srv.Image, srv.Plans.Specs.Cpus.Frequency, i.Address, k.Href, srv.Region.Name, srv.State)
				}
			}

		}
	}
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----------\t--------\t------\t-----\n")
	tw.Flush()

}

func listServer(c *cherrygo.Client, serverID string) {

	server, _, err := c.Server.List(serverID)
	if err != nil {
		log.Fatalf("Error while listing server: %v", err)
	}

	srvPower, _, err := c.Server.PowerState(serverID)
	if err != nil {
		log.Fatalf("Error while getting power sstate: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t---\t----------\t-----\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tPrice\tIP address\tPower\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----------\t-----\t-----\n")
	if len(server.IPAddresses) > 0 {
		for _, i := range server.IPAddresses {
			fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
				server.ID, server.Name, server.Hostname, server.Image, server.Pricing.Price, i.Address, srvPower.Power, server.State)
		}
	} else {
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			server.ID, server.Name, server.Hostname, server.Image, server.Pricing.Price, nil, srvPower.Power, server.State)
	}

	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----------\t-----\n")
	tw.Flush()

}

func createSSHKey(c *cherrygo.Client, label, key, keyPath string) {

	if keyPath != "" {
		keyData, err := ioutil.ReadFile(keyPath)
		if err != nil {
			log.Fatal("Error while reading key file: ", err)
		}

		if keyData != nil {
			key = string(keyData)
		}
	}

	sshCreateRequest := cherrygo.CreateSSHKey{
		Label: label,
		Key:   key,
	}

	sshkey, _, err := c.SSHKey.Create(&sshCreateRequest)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n------\t-----\t-----------\t-------\n")
	fmt.Fprintf(tw, "KEY ID\tLabel\tFingerprint\tCreated\n")
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\n", sshkey.ID, sshkey.Label, sshkey.Fingerprint, sshkey.Created)
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\n")
	tw.Flush()
}

func listSSHkeys(c *cherrygo.Client) {

	sshkeys, _, err := c.SSHKeys.List()
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n------\t-----\t-----------\t-------\t-------\n")
	fmt.Fprintf(tw, "Key ID\tLabel\tFingerprint\tCreated\tUpdated\n")
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\t-------\n")
	for _, s := range sshkeys {
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n",
			s.ID, s.Label, s.Fingerprint, s.Created, s.Updated)
	}
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\t-------\n")
	tw.Flush()
}

func listSSHkey(c *cherrygo.Client, keyID string) {

	sshkey, _, err := c.SSHKey.List(keyID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n------\t-----\t-----------\t-------\t-------\n")
	fmt.Fprintf(tw, "Key ID\tLabel\tFingerprint\tCreated\tUpdated\n")
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\t-------\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n",
		sshkey.ID, sshkey.Label, sshkey.Fingerprint, sshkey.Created, sshkey.Updated)
	fmt.Fprintf(tw, "------\t-----\t-----------\t-------\t-------\n")
	tw.Flush()
}

func updateSSHKey(c *cherrygo.Client, keyID, keyLabel string) {

	sshUpateRequest := cherrygo.UpdateSSHKey{
		Label: keyLabel,
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t-----\t-----------\t-------\t-------\n")
	fmt.Fprintf(tw, "Key ID\tLabel\tFingerprint\tCreated\tUpdated\n")
	fmt.Fprintf(tw, "-------\t-----\t-----------\t-------\t-------\n")

	key, _, err := c.SSHKey.Update(keyID, &sshUpateRequest)
	if err != nil {
		log.Fatal("Error while updating SSH key", err)
	}

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n",
		key.ID, key.Label, key.Fingerprint, key.Created, key.Updated)
	fmt.Fprintf(tw, "-------\t-----\t-----------\t-------\t-------\n")
	tw.Flush()
}

func updateIPAddress(c *cherrygo.Client, ptrRecord, aRecord, routedTo, assignedTo, ipID, projectID string) {

	updateIPRequest := cherrygo.UpdateIPAddress{
		PtrRecord:  ptrRecord,
		ARecord:    aRecord,
		RoutedTo:   routedTo,
		AssignedTo: assignedTo,
	}

	ip, _, err := c.IPAddress.Update(projectID, ipID, &updateIPRequest)
	if err != nil {
		log.Fatalf("Error while updating floating IP address: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "IP ID\tAddress\tCidr\tPTR\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n", ip.ID, ip.Address, ip.Cidr, ip.PtrRecord, ip.Type)
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t----\n")
	tw.Flush()
}

func deleteSSHKey(c *cherrygo.Client, keyID string) {

	sshDeleteRequest := cherrygo.DeleteSSHKey{ID: keyID}

	c.SSHKey.Delete(&sshDeleteRequest)
}

func deleteServer(c *cherrygo.Client, serverID string) {

	serverDeleteRequest := cherrygo.DeleteServer{ID: serverID}

	c.Server.Delete(&serverDeleteRequest)
}

func deleteIPAddress(c *cherrygo.Client, projectID, ipID string) {

	log.Println("DELETE IP ADDRESS")
	ipDeleteRequest := cherrygo.RemoveIPAddress{ID: ipID}

	c.IPAddress.Remove(projectID, &ipDeleteRequest)
}

func listPlans(c *cherrygo.Client, projectID int) {

	plans, _, err := c.Plans.List(projectID)
	if err != nil {
		log.Fatalf("Plans error: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t-------------\t----------\t-------------\t---\t-------\n")
	fmt.Fprintf(tw, "Plan ID\tPlan name\tPlan price\tCPU\tRAM\tRegions\n")
	fmt.Fprintf(tw, "-------\t-------------\t----------\t-------------\t---\t-------\n")
	for _, p := range plans {

		var regions []string
		for _, r := range p.AvailableRegions {
			regions = append(regions, r.Name)
		}
		reg := strings.Join(regions, ",")

		var planPrice float32
		for _, price := range p.Pricing {
			if price.Unit == "Hourly" {
				planPrice = price.Price
				// Prints most expensive plans
				if planPrice > 0 {
					fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n",
						p.ID, p.Name, planPrice, p.Specs.Cpus.Name, p.Specs.Memory.Total, reg)

				}
			}
		}
	}
	fmt.Fprintf(tw, "-------\t-------------\t----------\t-------------\t---\t-------\n")
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

func listProjects(c *cherrygo.Client, teamID int) {

	// Needs teams id to be passed
	projects, _, err := c.Projects.List(teamID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n----------\t------------\t----\n")
	fmt.Fprintf(tw, "Project ID\tProject name\tHref\n")
	fmt.Fprintf(tw, "----------\t------------\t----\n")

	for _, p := range projects {
		fmt.Fprintf(tw, "%v\t%v\t%v\n",
			p.ID, p.Name, p.Href)
	}
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	tw.Flush()
}

func powerOn(c *cherrygo.Client, serverID string) {

	server, _, err := c.Server.PowerOn(serverID)
	if err != nil {
		log.Fatal("Error while powering on the server", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t-----\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tPrice\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n",
		server.ID, server.Name, server.Hostname, server.Image, server.Pricing.Price, server.State)

	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")
	tw.Flush()

}

func powerOff(c *cherrygo.Client, serverID string) {

	server, _, err := c.Server.PowerOff(serverID)
	if err != nil {
		log.Fatal("Error while powering on the server", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t-----\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tPrice\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n",
		server.ID, server.Name, server.Hostname, server.Image, server.Pricing.Price, server.State)

	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")
	tw.Flush()
}

func rebootServer(c *cherrygo.Client, serverID string) {

	server, _, err := c.Server.Reboot(serverID)
	if err != nil {
		log.Fatal("Error while rebooting the server", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t-----\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tPrice\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n",
		server.ID, server.Name, server.Hostname, server.Image, server.Pricing.Price, server.State)

	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t-----\t-----\n")
	tw.Flush()
}

func listIPAddresses(c *cherrygo.Client, projectID string) {

	ips, _, err := c.IPAddresses.List(projectID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t---------\t----\n")
	fmt.Fprintf(tw, "IP ID\tAddress\tCidr\tPTR\tRouted To\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t---------\t----\n")
	for _, ip := range ips {
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n", ip.ID, ip.Address, ip.Cidr, ip.PtrRecord, ip.RoutedTo.Address, ip.Type)
	}
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t---------\t----\n")
	tw.Flush()
}

func listIPAddress(c *cherrygo.Client, projectID, ipID string) {

	ipp, _, err := c.IPAddress.List(projectID, ipID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t---------\t----\n")
	fmt.Fprintf(tw, "IP ID\tAddress\tCidr\tPTR\tRouted To\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t---------\t----\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n", ipp.ID, ipp.Address, ipp.Cidr, ipp.PtrRecord, ipp.RoutedTo.Address, ipp.Type)

	fmt.Fprintf(tw, "-----\t-------\t----\t---\t---------\t----\n")
	tw.Flush()
}

func printDebugInfo(c *cherrygo.Client) {

	log.Println(c.AuthToken)
	log.Println(c.BaseURL)
	log.Println(c.UserAgent)

}
