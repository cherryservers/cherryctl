package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

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