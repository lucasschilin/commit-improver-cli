package git

import (
	"os/exec"
	"strings"
)

func GetRepoRoot() (string, error) {

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
