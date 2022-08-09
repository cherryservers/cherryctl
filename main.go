/*
Copyright © 2022 Cherry Severs <support@cherryservers.com>

*/
package main

import (
	"os"

	"github.com/cherryservers/cherryctl/cmd"
)

func main() {
	cli := cmd.NewCli()
	if err := cli.MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
