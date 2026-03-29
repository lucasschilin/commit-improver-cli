package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var hookUninstallCmd = &cobra.Command{
	Use:   "hook-uninstall",
	Short: "Uninstall git hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}

		err = git.RemoveCommitMsgHook(repoRoot)
		if err != nil {
			return fmt.Errorf("Error uninstalling cim-cli hook: %v", err)
		}

		fmt.Println("cim-cli hook uninstalled :(")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hookUninstallCmd)
}
