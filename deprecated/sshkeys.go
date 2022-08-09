package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	"github.com/cherryservers/cherrygo"
)

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

func deleteSSHKey(c *cherrygo.Client, keyID string) {

	sshDeleteRequest := cherrygo.DeleteSSHKey{ID: keyID}

	c.SSHKey.Delete(&sshDeleteRequest)
}
