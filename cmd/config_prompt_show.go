package cmd

import (
	"errors"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/lucasschilin/cim-cli/internal/prompt"
	"github.com/spf13/cobra"
)

var configPromptShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show prompt",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			return errors.New("Not inside a git repository")
		}

		var finalPrompt string
		var found bool
		var err error

		switch {
		case globalFlag:
			finalPrompt, found, err = prompt.Load(prompt.GlobalPromptPath())

		case repoUserFlag:
			finalPrompt, found, err = prompt.Load(prompt.RepoUserPromptPath(repoRoot))

		case repoFlag:
			finalPrompt, found, err = prompt.Load(prompt.RepoSharedPromptPath(repoRoot))

		default:
			finalPrompt, err = prompt.ResolveTemplate(repoRoot)
		}

		if err != nil {
			return fmt.Errorf("Prompt config error: %v", err)
		} else if !found && (globalFlag || repoUserFlag || repoFlag) {
			return errors.New("Prompt config not found")
		}

		fmt.Println(finalPrompt)

		return nil

	},
}

func init() {
	configPromptCmd.AddCommand(configPromptShowCmd)

	configPromptShowCmd.Flags().BoolVar(&globalFlag, "global", false, "Show global prompt")
	configPromptShowCmd.Flags().BoolVar(&repoFlag, "repo", false, "Show shared repo prompt")
	configPromptShowCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Show user prompt for this repo")
}
