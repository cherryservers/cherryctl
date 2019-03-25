package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

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
