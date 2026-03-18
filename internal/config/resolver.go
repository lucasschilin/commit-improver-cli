package config

import (
	"os"
)

// ResolveConfig carrega e combina as configurações dos 3 níveis.
func Resolve(repoRoot string) (*Config, error) {

	final := &Config{}

	paths := []string{GlobalConfigPath()}

	if repoRoot != "" {
		paths = append(paths, RepoUserConfigPath(repoRoot))
		paths = append(paths, RepoSharedConfigPath(repoRoot))
	}

	for _, path := range paths {

		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		cfg, err := LoadConfigFile(path)
		if err != nil {
			return nil, err
		}

		mergeConfig(final, cfg)
	}

	final.ApplyDefaults()

	if err := final.Validate(); err != nil {
		return nil, err
	}

	return final, nil
}

func mergeConfig(dst *Config, src *Config) {

	if src.Language != "" {
		dst.Language = src.Language
	}

	if src.DiffLimit != 0 {
		dst.DiffLimit = src.DiffLimit
	}

	if src.ImprovementRequestTimeout != 0 {
		dst.ImprovementRequestTimeout = src.ImprovementRequestTimeout
	}

	if src.AllowFinalEdit == true {
		dst.AllowFinalEdit = src.AllowFinalEdit
	}

	if src.Provider != "" {
		dst.Provider = src.Provider
	}

	if src.Model != "" {
		dst.Model = src.Model
	}

	if src.Gemini.APIKey != "" {
		dst.Gemini.APIKey = src.Gemini.APIKey
	}
}
