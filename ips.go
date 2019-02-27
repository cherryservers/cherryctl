package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
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

func deleteIPAddress(c *cherrygo.Client, projectID, ipID string) {

	log.Println("DELETE IP ADDRESS")
	ipDeleteRequest := cherrygo.RemoveIPAddress{ID: ipID}

	c.IPAddress.Remove(projectID, &ipDeleteRequest)
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
