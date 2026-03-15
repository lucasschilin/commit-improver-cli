package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasschilin/cim-cli/internal/ai"
	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/lucasschilin/cim-cli/internal/prompt"
	"github.com/lucasschilin/cim-cli/internal/ui"
	"github.com/spf13/cobra"
)

var messageFlag string

var improveCmd = &cobra.Command{
	Use:   "improve",
	Short: "Improve commit message",
	Run: func(cmd *cobra.Command, args []string) {
		var message string

		if messageFlag != "" {
			message = messageFlag
		}

		if message == "" {
			fmt.Println("Commit message not provided")
			return
		}

		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			fmt.Println("Not inside a git repository")
		}
		cfg, err := config.Resolve(repoRoot)
		if err != nil {
			fmt.Println("Config error:", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		aiCfg := ai.Config{
			Provider: cfg.Provider,
			Model:    cfg.Model,
			APIKey:   cfg.Gemini.APIKey,
		}

		provider, err := ai.NewProvider(ctx, aiCfg)
		if err != nil {
			fmt.Println("Error creating AI provider:", err)
			return
		}

		prompt := prompt.Build(message, "", cfg.Language)

		spinner := ui.New("Improving commit message...\n")
		spinner.Start()
		defer spinner.Stop()

		improvedMessage, err := provider.ImproveCommitMessage(ctx, prompt)
		if err != nil {
			spinner.Stop()
			fmt.Println("✖ Failed to improve commit:", err)
			return
		}

		spinner.Stop()
		fmt.Print("✔ Commit message improved\n\n")

		ui.ShowPreview(message, improvedMessage)

	},
}

func init() {
	rootCmd.AddCommand(improveCmd)

	improveCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Commit message to improve")
}
