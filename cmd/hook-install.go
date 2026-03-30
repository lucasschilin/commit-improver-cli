package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var hookInstallCmd = &cobra.Command{
	Use:   "hook-install",
	Short: "Install git hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}
		_, err = config.Resolve(repoRoot)
		if err != nil {
			return fmt.Errorf("Config error: %v", err)
		}

		err = git.InstallCommitMsgHook(repoRoot)
		if err != nil {
			return fmt.Errorf("Error installing cim-cli hook: %v", err)

		}

		fmt.Println("\u2705 cim-cli hook INSTALLED successfully")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hookInstallCmd)
}
