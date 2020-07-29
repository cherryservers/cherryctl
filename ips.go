package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

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

func updateIPAddress(c *cherrygo.Client, ptrRecord, aRecord, routedTo,
	routedToHostname, routedToServerIP, routedToServerID, assignedTo, floatingID, floatingIP, projectID string) {

	if routedToHostname != "" || routedToServerIP != "" || routedToServerID != "" {
		res, err := getIDForServerIP(c, projectID, routedToHostname, routedToServerIP, routedToServerID)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		routedTo = res
	}

	updateIPRequest := cherrygo.UpdateIPAddress{
		PtrRecord:  ptrRecord,
		ARecord:    aRecord,
		RoutedTo:   routedTo,
		AssignedTo: assignedTo,
	}

	if floatingIP != "" || floatingID != "" {
		flID, err := getIDForFloatingIP(c, projectID, floatingIP, floatingID)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		floatingID = flID
	}

	ip, _, err := c.IPAddress.Update(projectID, floatingID, &updateIPRequest)
	if err != nil {
		log.Fatalf("Error while updating floating IP address: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n-----\t-------\t----\t---\t-\t----\n")
	fmt.Fprintf(tw, "IP ID\tAddress\tCidr\tPTR\tA\tType\n")
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t-\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\t%v\n", ip.ID, ip.Address, ip.Cidr, ip.PtrRecord, ip.ARecord, ip.Type)
	fmt.Fprintf(tw, "-----\t-------\t----\t---\t-\t----\n")
	tw.Flush()
}

func deleteIPAddress(c *cherrygo.Client, projectID, ipID string) {

	log.Println("DELETE IP ADDRESS")
	ipDeleteRequest := cherrygo.RemoveIPAddress{ID: ipID}

	c.IPAddress.Remove(projectID, &ipDeleteRequest)
}

func listIPAddresses(c *cherrygo.Client, projectID string) {

	ips, _, err := c.IPAddresses.List(projectID)
	if err != nil {
		log.Fatalf("Error while listing IP addresses: %v", err)
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

func getIDForFloatingIP(c *cherrygo.Client, projectID, floatingIP, floatingID string) (string, error) {

	var flID string

	ips, _, err := c.IPAddresses.List(projectID)
	if err != nil {
		err = fmt.Errorf("something went wrong while listing IP addresses - %v", err)
	}

	for _, ip := range ips {
		switch {
		case floatingIP != "":
			if ip.Address == floatingIP {
				flID = ip.ID
			}
		case floatingID != "":
			if ip.ID == floatingID {
				flID = ip.ID
			}
		}
	}

	if floatingIP == "" && flID == "" {
		err = fmt.Errorf("it seems there is no such floating ID - %v", floatingID)
	}

	if floatingID == "" && flID == "" {
		err = fmt.Errorf("it seems there is no such floating IP - %v", floatingIP)
	}

	return flID, err
}

func getIDForServerIP(c *cherrygo.Client, projectID, routedToHostname, routedToServerIP, routedToServerID string) (string, error) {

	servers, _, err := c.Servers.List(projectID)
	if err != nil {
		log.Fatalf("Error while listing server: %v", err)
	}

	var routeTo string
	var item string
	var sliceOfkeys []string

	myDict := make(map[string]string)
	uniqDict := make(map[string]string)
	nonUniqDict := make(map[string]string)

	switch {
	case routedToHostname != "":
		item = routedToHostname
		for _, srv := range servers {

			serverID := strconv.Itoa(srv.ID)
			if srv.Hostname == routedToHostname {
				myDict[serverID] = srv.Hostname
				if len(srv.IPAddresses) > 0 {
					for _, i := range srv.IPAddresses {
						if i.Type == "primary-ip" {
							routeTo = i.ID
						}
					}
				}
			}
		}
	case routedToServerIP != "":
		item = routedToServerIP
		for _, srv := range servers {
			serverID := strconv.Itoa(srv.ID)
			if len(srv.IPAddresses) > 0 {
				for _, i := range srv.IPAddresses {
					if i.Type == "primary-ip" {
						if i.Address == routedToServerIP {
							myDict[serverID] = i.Address
							routeTo = i.ID
						}
					}
				}
			}
		}
	case routedToServerID != "":
		item = routedToServerID
		for _, srv := range servers {
			serverID := strconv.Itoa(srv.ID)
			if serverID == routedToServerID {
				myDict[serverID] = serverID
				if len(srv.IPAddresses) > 0 {
					for _, i := range srv.IPAddresses {
						if i.Type == "primary-ip" {
							routeTo = i.ID
						}
					}
				}
			}
		}
	}

	for k, v := range myDict {
		if v == item {
			uniqDict[k] = v
			sliceOfkeys = append(sliceOfkeys, k)
		}
	}

	if len(uniqDict) == 0 {
		err = fmt.Errorf("it seems item %v can't be found. Please check it and try again", item)
	}

	if len(sliceOfkeys) > 1 {
		for _, v := range sliceOfkeys {
			nonUniqDict[v] = item
		}
		err = fmt.Errorf("there are several nodes with same hostname: %v. Please use routed_to_server_id instead", nonUniqDict)
	}

	return routeTo, err
}
