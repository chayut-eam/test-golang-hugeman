package meta

import (
	"bufio"
	"strings"

	log "github.com/sirupsen/logrus"
)

type InfoResponse struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	BuiltTimestamp  string `json:"built_timestamp"`
	Commit          string `json:"commit"`
	CommitTimestamp string `json:"commit_timestamp"`
}

func ParseInfo(resource string) InfoResponse {
	log := log.StandardLogger()
	log.Info("Parsing info.properties")

	properties := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(resource))

	// parse each line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip commented line
		if !strings.HasPrefix(line, "#") {
			prop := strings.Split(line, "=")
			if len(prop) >= 2 {
				properties[prop[0]] = prop[1]
			} else {
				log.Warn("Skip invalid property, ", line)
			}
		}
	}

	return InfoResponse{
		Name:            properties["name"],
		Version:         properties["version"],
		BuiltTimestamp:  properties["built_timestamp"],
		Commit:          properties["commit"],
		CommitTimestamp: properties["commit_timestamp"],
	}
}
