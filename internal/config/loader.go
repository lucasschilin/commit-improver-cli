package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

// LoadConfigFile lê e valida um arquivo YAML.
func LoadConfigFile(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid yaml in %s: %w", path, err)
	}

	return cfg, nil
}
