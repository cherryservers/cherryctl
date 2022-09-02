package utils

import (
	"fmt"
	"strings"

	"github.com/cherryservers/cherrygo/v3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func BoolToYesNo(b bool) string {
	if b {
		return "yes"
	}

	return "no"
}

func ServerHostnameToID(hostname string, projectID int, ServerService cherrygo.ServersService) (int, error) {
	if projectID == 0 {
		return 0, fmt.Errorf("Project ID must be set")
	}

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

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
