package config

type Config struct {
	Provider  string `yaml:"provider"`
	Model     string `yaml:"model"`
	Language  string `yaml:"language"`
	DiffLimit int    `yaml:"diff_limit"`

	Gemini struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"gemini"`
}
