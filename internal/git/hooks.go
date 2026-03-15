package git

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CommitMsgHookPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".git", "hooks", "commit-msg")
}

func HookExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsCommitImproverCLIHook(content string) bool {
	return strings.Contains(content, "cim-cli hook")
}

func InstallCommitMsgHook(repoRoot string) error {

	path := CommitMsgHookPath(repoRoot)

	if HookExists(path) {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if IsCommitImproverCLIHook(string(data)) {
			return fmt.Errorf("cim-cli hook already installed")
		}
		return fmt.Errorf("another commit-msg hook already exists")
	}

	hook := `#!/bin/sh
# cim-cli hook

cim-cli hook "$1"

`

	err := os.WriteFile(path, []byte(hook), 0755)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCommitMsgHook(repoRoot string) error {
	path := CommitMsgHookPath(repoRoot)

	if !HookExists(path) {
		return fmt.Errorf("cim-cli hook not installed")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if !IsCommitImproverCLIHook(string(data)) {
		return fmt.Errorf("commit-msg is not managed by cim-cli")
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
