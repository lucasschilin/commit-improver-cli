package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var deinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "Remove cim-cli git hook",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			fmt.Println("Not inside a git repository")
			return
		}

		err = git.RemoveCommitMsgHook(repoRoot)
		if err != nil {
			fmt.Println("Error removing cim-cli hook:", err)
			return
		}

		fmt.Println("cim-cli hook removed :(")
	},
}

func init() {
	rootCmd.AddCommand(deinitCmd)
}
