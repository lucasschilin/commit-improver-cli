package config

import "errors"

type Config struct {
	Language                  string `yaml:"language"`
	DiffLimit                 *int   `yaml:"diff_limit"`
	AllowFinalEdit            bool   `yaml:"allow_final_edit"`
	ImprovementRequestTimeout *int   `yaml:"improvement_request_timeout"`

	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
	Gemini   struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"gemini"`
	Openai struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"openai"`
}

func (c *Config) ApplyDefaults() {
	if c.Language == "" {
		c.Language = "en"
	}

	if c.DiffLimit == nil {
		val := 0
		c.DiffLimit = &val
	}

	if c.ImprovementRequestTimeout == nil {
		val := 20
		c.ImprovementRequestTimeout = &val
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

	if c.DiffLimit != nil && *c.DiffLimit < 0 {
		return errors.New("config: diff_limit must be 0 or greater")
	}

	if c.ImprovementRequestTimeout != nil && *c.ImprovementRequestTimeout <= 0 {
		return errors.New("config: improvement_request_timeout is required and must be greater than zero")
	}

	if c.Provider == "" {
		return errors.New("config: provider is required")
	}

	if c.Model == "" {
		return errors.New("config: model is required")
	}

	switch c.Provider {
	case "gemini":
		if c.Gemini.APIKey == "" {
			return errors.New("config: gemini.api_key is required")
		}
	case "openai":
		if c.Openai.APIKey == "" {
			return errors.New("config: openai.api_key is required")
		}
	}

	return nil
}
