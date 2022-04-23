package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

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
