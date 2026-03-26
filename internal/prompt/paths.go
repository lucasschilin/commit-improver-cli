package prompt

import (
	"os"
	"path/filepath"
)

func GlobalPromptPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".cim-cli", "prompt.md")
}

func RepoSharedPromptPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".cim-cli", "prompt.md")
}

func RepoUserPromptPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".git", ".cim-cli", "prompt.md")
}
