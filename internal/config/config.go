package config

import "errors"

type Config struct {
	Language                  string `yaml:"language"`
	DiffLimit                 int    `yaml:"diff_limit"`
	AllowFinalEdit            bool   `yaml:"allow_final_edit"`
	ImprovementRequestTimeout int    `yaml:"improvement_request_timeout"`

	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
	Gemini   struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"gemini"`
}

func (c *Config) ApplyDefaults() {
	if c.Language == "" {
		c.Language = "en"
	}

	if c.DiffLimit == 0 {
		c.DiffLimit = 200
	}

	if c.ImprovementRequestTimeout == 0 {
		c.ImprovementRequestTimeout = 20
	}

	if c.Provider == "" {
		c.Provider = "gemini"
	}

	if c.Model == "" {
		c.Model = "gemini-2.5-flash"
	}
}

func (c *Config) Validate() error {
	if c.Language == "" {
		return errors.New("config: language is required")
	}

	if c.DiffLimit <= 0 {
		return errors.New("config: diff_limit is required and must be greater than zero")
	}

	if c.ImprovementRequestTimeout <= 0 {
		return errors.New("config: improvement_request_timeout is required and must be greater than zero")
	}

	if c.Provider == "" {
		return errors.New("config: provider is required")
	}

	if c.Model == "" {
		return errors.New("config: model is required")
	}

	if c.Provider == "gemini" {
		if c.Gemini.APIKey == "" {
			return errors.New("config: gemini.api_key is required")
		}
	}

	return nil
}
