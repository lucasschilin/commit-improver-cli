package cmd

import (
	"fmt"
	"os"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var hookStatusCmd = &cobra.Command{
	Use:   "hook-status",
	Short: "View git hook status",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}
		_, err = config.Resolve(repoRoot)
		if err != nil {
			return fmt.Errorf("Config error: %v", err)
		}

		path := git.CommitMsgHookPath(repoRoot)
		if git.HookExists(path) {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if git.IsCommitImproverCLIHook(string(data)) {
				fmt.Println("cim-cli hook installed :D")
				return nil
			}

			fmt.Println("Another commit-msg hook installed :(")
			return nil
		}

		fmt.Println("cim-cli hook not installed. Run the hook install command to install.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hookStatusCmd)
}
