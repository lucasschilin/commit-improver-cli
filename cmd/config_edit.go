package cmd

import (
	"errors"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/editor"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			return errors.New("Not inside a git repository")
		}

		var path string

		switch {
		case globalFlag:
			path = config.GlobalConfigPath()

		case repoUserFlag:
			path = config.RepoUserConfigPath(repoRoot)

		case repoFlag:
			path = config.RepoSharedConfigPath(repoRoot)

		default:
			return errors.New("You need to inform which configuration to edit.\nUse one of the following flags: --global, --repo, --repo-user")
		}

		err := config.EnsureConfigFile(path)
		if err != nil {
			return fmt.Errorf("Error creating config file: %v", err)
		}

		editor.Open(path)

		return nil

	},
}

func init() {
	configCmd.AddCommand(configEditCmd)

	configEditCmd.Flags().BoolVar(&globalFlag, "global", false, "Edit global config")
	configEditCmd.Flags().BoolVar(&repoFlag, "repo", false, "Edit shared repo config")
	configEditCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Edit user config for this repo")
}
