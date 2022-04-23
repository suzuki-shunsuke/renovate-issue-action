package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RenovateLogin string `yaml:"renovate_login"`
}

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

func Read(p string, cfg *Config) error {
	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("open a configuration file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("parse configuration file as YAML: %w", err)
	}
	return nil
}

func SetDefault(cfg *Config) {
	if cfg.RenovateLogin == "" {
		cfg.RenovateLogin = "renovate[bot]"
	}
}
