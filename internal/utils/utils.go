package utils

import (
	"bytes"
	"fmt"
	"os"
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

func FormatStringTags(tags *map[string]string) string {
	var buffer bytes.Buffer

	for key, val := range *tags {
		buffer.WriteString(key + ":" + val + " ")
	}

	return buffer.String()
}

// ReadOptionalFile is identical to os.ReadFile, except that if 'path'
// is empty, it returns a nil slice, with no error.
func ReadOptionalFile(path string) ([]byte, error) {
	if path == "" {
		return nil, nil
	}

	return os.ReadFile(path)
}
