package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func addServer(c *cherrygo.Client, projectID, hostname string, ipaddresses, sshKeys []string, image, planID, region, userData string) {

	addServerRequest := cherrygo.CreateServer{
		ProjectID:   projectID,
		Hostname:    hostname,
		Image:       image,
		Region:      region,
		SSHKeys:     sshKeys,
		IPAddresses: ipaddresses,
		PlanID:      planID,
		UserData:    userData,
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
	fmt.Fprintf(tw, "\n---------\t----\t--------\t-----\t---\t----\t----------\t--------\t------\t-----\n")
	fmt.Fprintf(tw, "Server ID\tName\tHostname\tImage\tCPU2\tUnit\tIP address\tSSH keys\tRegion\tState\n")
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----\t----------\t--------\t------\t-----\n")

	for _, srv := range servers {

		for _, i := range srv.IPAddresses {
			if len(srv.SSHKeys) == 0 {
				fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
					srv.ID, srv.Name, srv.Hostname, srv.Image, srv.Plans.Specs.Cpus.Frequency, srv.Plans.Specs.Cpus.Unit, i.Address, nil, srv.Region.Name, srv.State)
			} else {
				for _, k := range srv.SSHKeys {
					fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
						srv.ID, srv.Name, srv.Hostname, srv.Image, srv.Plans.Specs.Cpus.Frequency, srv.Plans.Specs.Cpus.Unit, i.Address, k.Href, srv.Region.Name, srv.State)
				}
			}

		}
	}
	fmt.Fprintf(tw, "---------\t----\t--------\t-----\t---\t----\t----------\t--------\t------\t-----\n")
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

func deleteServer(c *cherrygo.Client, serverID string) {

	serverDeleteRequest := cherrygo.DeleteServer{ID: serverID}

	c.Server.Delete(&serverDeleteRequest)
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
