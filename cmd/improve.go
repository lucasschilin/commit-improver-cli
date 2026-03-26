package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lucasschilin/cim-cli/internal/ai"
	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/editor"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/lucasschilin/cim-cli/internal/prompt"
	"github.com/lucasschilin/cim-cli/internal/ui"
	"github.com/spf13/cobra"
)

var (
	messageFlag string
	promptFlag  bool
)

var improveCmd = &cobra.Command{
	Use:   "improve",
	Short: "Improve commit message",
	RunE: func(cmd *cobra.Command, args []string) error {
		var message string

		if messageFlag != "" {
			message = messageFlag
		}

		if message == "" {
			msg, err := editor.OpenTempFile()
			if err != nil {
				return fmt.Errorf("Failed to open editor: %v", err)
			}

			message = strings.TrimSpace(msg)
		}

		if message == "" {
			return errors.New("Commit message not provided")
		}

		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}
		cfg, err := config.Resolve(repoRoot)
		if err != nil {
			return fmt.Errorf("Config error: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*cfg.ImprovementRequestTimeout)*time.Second)
		defer cancel()

		customPrompt, found, err := prompt.Resolve(repoRoot) // TODO: use here diff and message
		if err != nil {
			return fmt.Errorf("prompt error: %v", err)
		}

		var finalPrompt string

		if found {
			finalPrompt = customPrompt
		} else {
			finalPrompt = prompt.Build(message, "", cfg.Language)
		}

		if promptFlag {
			fmt.Println("=== GENERATED PROMPT ===")
			fmt.Println(finalPrompt)
			return nil
		}

		provider, err := ai.NewProvider(ctx, cfg)
		if err != nil {
			return fmt.Errorf("Error creating AI provider: %v", err)
		}

		spinner := ui.New("Improving commit message...\n")
		spinner.Start()
		defer spinner.Stop()

		improvedMessage, err := provider.ImproveCommitMessage(ctx, finalPrompt)
		if err != nil {
			spinner.Stop()
			return fmt.Errorf("✖ Failed to improve commit: %v", err)
		}

		spinner.Stop()
		fmt.Print("✔ Commit message improved\n\n")

		ui.ShowPreview(message, improvedMessage)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(improveCmd)

	improveCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Commit message to improve")
	improveCmd.Flags().BoolVarP(&promptFlag, "prompt", "p", false, "Prints the final message that would be sent to the LLM.")
}
