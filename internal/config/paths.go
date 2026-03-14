package config

import (
	"os"
	"path/filepath"
)

func GlobalConfigPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".commit-improver-cli", "config.yaml")
}

func RepoSharedConfigPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".commit-improver-cli", "config.yaml")
}

func RepoUserConfigPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".git", ".commit-improver-cli", "config.yaml")
}
