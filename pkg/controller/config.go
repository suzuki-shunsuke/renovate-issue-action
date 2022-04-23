package controller

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RenovateLogin string `yaml:"renovate_login"`
}

var defaultConfigPaths []string = []string{
	"renovate-issue-action.yaml",
	"renovate-issue-action.yml",
	".renovate-issue-action.yaml",
	".renovate-issue-action.yml",
}

func findConfig(p string) string {
	if p != "" {
		return p
	}
	for _, p := range defaultConfigPaths {
		if _, err := os.Stat(p); err != nil {
			return p
		}
	}
	return ""
}

func readConfig(p string, cfg *Config) error {
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

func setDefaultConfig(cfg *Config) {
	if cfg.RenovateLogin == "" {
		cfg.RenovateLogin = "renovate[bot]"
	}
}
