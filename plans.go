package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

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
