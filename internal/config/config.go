package config

import "errors"

type Config struct {
	Provider  string `yaml:"provider"`
	Model     string `yaml:"model"`
	Language  string `yaml:"language"`
	DiffLimit int    `yaml:"diff_limit"`

	Gemini struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"gemini"`
}

func (c *Config) ApplyDefaults() {
	if c.Provider == "" {
		c.Provider = "gemini"
	}

	if c.Model == "" {
		c.Model = "gemini-2.5-flash"
	}

	if c.Language == "" {
		c.Language = "en"
	}

	if c.DiffLimit == 0 {
		c.DiffLimit = 200
	}
}

func (c *Config) Validate() error {

	if c.Provider == "" {
		return errors.New("config: provider is required")
	}

	if c.Model == "" {
		return errors.New("config: model is required")
	}

	if c.Language == "" {
		return errors.New("config: language is required")
	}

	if c.DiffLimit <= 0 {
		return errors.New("config: diff_limit is required and must be greater than zero")
	}

	if c.Provider == "gemini" && c.Gemini.APIKey == "" {
		return errors.New("config: gemini.api_key is required")
	}

	return nil
}
