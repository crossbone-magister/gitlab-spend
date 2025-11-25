package config

import (
	"fmt"

	"github.com/crossbone-magister/timewlib"
)

const extensionPrefix = "reports.gitlabspend"
const hostConfigKey = extensionPrefix + ".host"
const tokenConfigKey = extensionPrefix + ".token"

type GitlabSpendConfiguration map[string]string

func New(configuration timewlib.Configuration) (GitlabSpendConfiguration, error) {
	var config GitlabSpendConfiguration = configuration.GetAllByPrefix(extensionPrefix)
	if config.Host() == "" {
		return nil, fmt.Errorf("no gitlab host configured at %s", hostConfigKey)
	}
	if config.Token() == "" {
		return nil, fmt.Errorf("no gitlab token configured at %s", tokenConfigKey)
	}
	return config, nil
}

func (c GitlabSpendConfiguration) Host() string {
	return c[hostConfigKey]
}

func (c GitlabSpendConfiguration) Token() string {
	return c[tokenConfigKey]
}
