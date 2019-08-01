package main

import (
	"cherrygo"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

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
		var imagePrice float32

		if len(i.Pricing) > 0 {
			for _, price := range i.Pricing {

				if price.Unit == "Hourly" {
					imagePrice = price.Price
					fmt.Fprintf(tw, "%v\t%v\t%v\n",
						i.ID, i.Name, imagePrice)
				}

			}
		} else {
			fmt.Fprintf(tw, "%v\t%v\t%v\n",
				i.ID, i.Name, imagePrice)
		}
	}

	fmt.Fprintf(tw, "--------\t----------\t-----------\n")
	tw.Flush()
}
