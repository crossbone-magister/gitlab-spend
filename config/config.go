package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/crossbone-magister/timewlib"
)

const extensionPrefix = "reports.gitlabspend"
const hostConfigKey = extensionPrefix + ".host"
const tokenConfigKey = extensionPrefix + ".token"
const stateFileConfigKey = extensionPrefix + ".state_file"

type GitlabSpendConfiguration map[string]string

func New(configuration timewlib.Configuration) (GitlabSpendConfiguration, error) {
	var config GitlabSpendConfiguration = configuration.GetAllByPrefix(extensionPrefix)
	// Timewarrior only passes its built-in keys via stdin; read custom extension
	// config directly from the config file using the temp.config path.
	if cfgPath := configuration["temp.config"]; cfgPath != "" {
		if fileConfig, err := parseConfigFile(cfgPath); err == nil {
			for k, v := range fileConfig {
				if _, exists := config[k]; !exists {
					config[k] = v
				}
			}
		}
	}
	if config.Host() == "" {
		return nil, fmt.Errorf("no gitlab host configured at %s", hostConfigKey)
	}
	if config.Token() == "" {
		return nil, fmt.Errorf("no gitlab token configured at %s", tokenConfigKey)
	}
	return config, nil
}

func parseConfigFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	result := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, extensionPrefix) {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		result[key] = value
	}
	return result, scanner.Err()
}

func (c GitlabSpendConfiguration) Host() string {
	return c[hostConfigKey]
}

func (c GitlabSpendConfiguration) Token() string {
	return c[tokenConfigKey]
}

func (c GitlabSpendConfiguration) StateFile() string {
	if path := c[stateFileConfigKey]; path != "" {
		return path
	}
	base := os.Getenv("XDG_DATA_HOME")
	if base == "" {
		base = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}
	return filepath.Join(base, "gitlab-spend", "state.json")
}
