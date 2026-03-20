package cmd

import (
	"errors"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var configTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Show configuration file default template",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			return errors.New("Not inside a git repository")
		}

		template := config.DefaultConfigTemplate

		fmt.Println(template)

		return nil
	},
}

func init() {
	configCmd.AddCommand(configTemplateCmd)
}
