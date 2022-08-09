package utils

import (
	"fmt"
	"strings"

	"github.com/cherryservers/cherrygo"
	"github.com/pkg/errors"
)

func BoolToYesNo(b bool) string {
	if b {
		return "yes"
	}

	return "no"
}

func ServerHostnameToID(hostname string, projectID int, ServerService cherrygo.ServersService) (int, error) {
	serversList, err := serverList(projectID, ServerService)
	for _, s := range serversList {
		if strings.EqualFold(hostname, s.Hostname) {
			return s.ID, err
		}
	}

	return 0, errors.Wrap(err, fmt.Sprintf("Could not find server with `%s` hostname", hostname))
}

func serverList(projectID int, ServerService cherrygo.ServersService) ([]cherrygo.Server, error) {
	getOptions := cherrygo.GetOptions{
		Fields: []string{"id", "name", "hostname"},
	}
	serverList, _, err := ServerService.List(projectID, &getOptions)

	return serverList, err
}
