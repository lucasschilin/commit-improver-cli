package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucasschilin/commit-improver-cli/internal/ai"
	"github.com/lucasschilin/commit-improver-cli/internal/commit"
	"github.com/lucasschilin/commit-improver-cli/internal/git"
	"github.com/lucasschilin/commit-improver-cli/internal/prompt"
	"github.com/lucasschilin/commit-improver-cli/internal/ui"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Intercept commit messages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Commit message file not provided")
			return
		}

		path := args[0]

		message, err := commit.ReadCommitMessage(path)
		if err != nil {
			fmt.Println("Error reading commit message:", err)
			return
		}

		diff, err := git.GetStagedDiff()
		if err != nil {
			fmt.Println("Error reading diff:", err)
			return
		}

		diff = git.LimitDiff(diff, 200)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//temporary
		if err := godotenv.Load(); err != nil {
			log.Println("Any .env file finded.")
		}
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("API_KEY not set")
			return
		}

		cfg := ai.Config{
			Provider: "gemini",
			APIKey:   apiKey,
			Model:    "gemini-2.5-flash",
		}

		provider, err := ai.NewProvider(ctx, cfg)
		if err != nil {
			fmt.Println("Provider error:", err)
			return
		}

		prompt := prompt.Build(message, diff, "en")

		improvedMessage, err := provider.ImproveCommitMessage(ctx, prompt)
		if err != nil {
			fmt.Println("AI error:", err)
			return
		}

		ui.ShowPreview(message, improvedMessage)

		accepted, err := ui.Confirm("Apply improved commit message?")
		if err != nil {
			fmt.Println(err)
			return
		}

		if !accepted {
			return
		}

		err = commit.WriteCommitMessage(path, improvedMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
