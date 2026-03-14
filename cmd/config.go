package cmd

import (
	"fmt"

	"github.com/lucasschilin/commit-improver-cli/internal/config"
	"github.com/lucasschilin/commit-improver-cli/internal/git"
	"github.com/spf13/cobra"
)

var (
	globalFlag   bool
	repoFlag     bool
	repoUserFlag bool
	editorFlag   bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit configuration",
	RunE: func(cmd *cobra.Command, args []string) error {

		repoRoot, _ := git.GetRepoRoot()

		var path string

		switch {

		case globalFlag:
			path = config.GlobalConfigPath()

		case repoUserFlag:
			path = config.RepoUserConfigPath(repoRoot)

		case repoFlag:
			path = config.RepoSharedConfigPath(repoRoot)

		default:
			fmt.Println(
				"You need to inform which configuration to edit.\nUse one of the following flags: --global, --repo, --repo-user",
			)
			return nil
		}

		err := config.EnsureConfigFile(path)
		if err != nil {
			return err
		}

		if editorFlag != false {
			return config.OpenEditor(path)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVar(&globalFlag, "global", false, "Edit global config")
	configCmd.Flags().BoolVar(&repoFlag, "repo", false, "Edit shared repo config")
	configCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Edit user config for this repo")

	configCmd.Flags().BoolVar(&editorFlag, "editor", false, "Edit config using editor")
}
