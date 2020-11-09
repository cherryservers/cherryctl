package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/cherryservers/cherrygo"
)

func listProject(c *cherrygo.Client, projectID string) {

	project, _, err := c.Project.List(projectID)
	if err != nil {
		log.Fatal("Error", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n----------\t------------\t----\n")
	fmt.Fprintf(tw, "Project ID\tProject name\tHref\n")
	fmt.Fprintf(tw, "----------\t------------\t----\n")

	fmt.Fprintf(tw, "%v\t%v\t%v\n", project.ID, project.Name, project.Href)
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	tw.Flush()
}

func createProject(c *cherrygo.Client, teamID int, projectName string) {

	addProjectRequest := cherrygo.CreateProject{
		Name: projectName,
	}

	project, _, err := c.Project.Create(teamID, &addProjectRequest)
	if err != nil {
		log.Fatal("Error while creating new project: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n----------\t------------\t----\n")
	fmt.Fprintf(tw, "Project ID\tProject name\tHref\n")
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\n", project.ID, project.Name, project.Href)
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	tw.Flush()

}

func updateProject(c *cherrygo.Client, projectID string, projectName string) {

	updateProjectRequest := cherrygo.UpdateProject{
		Name: projectName,
	}

	project, _, err := c.Project.Update(projectID, &updateProjectRequest)
	if err != nil {
		log.Fatal("Error while updating new project: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 13, 8, 2, '\t', 0)
	fmt.Fprintf(tw, "\n----------\t------------\t----\n")
	fmt.Fprintf(tw, "Project ID\tProject name\tHref\n")
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	fmt.Fprintf(tw, "%v\t%v\t%v\n", project.ID, project.Name, project.Href)
	fmt.Fprintf(tw, "----------\t------------\t----\n")
	tw.Flush()

}

func deleteProject(c *cherrygo.Client, projectID string) {

	projectDeleteRequest := cherrygo.DeleteProject{ID: projectID}

	c.Project.Delete(projectID, &projectDeleteRequest)
}
