package config

import (
	"os"
)

var defaultConfigPaths = []string{ //nolint:gochecknoglobals
	"renovate-issue-action.yaml",
	"renovate-issue-action.yml",
	".renovate-issue-action.yaml",
	".renovate-issue-action.yml",
}

func Find(p string) string {
	if p != "" {
		return p
	}
	for _, p := range defaultConfigPaths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}
