package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lucasschilin/cim-cli/internal/ai"
	"github.com/lucasschilin/cim-cli/internal/commit"
	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/editor"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/lucasschilin/cim-cli/internal/prompt"
	"github.com/lucasschilin/cim-cli/internal/ui"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Intercept commit messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}
		cfg, err := config.Resolve(repoRoot)
		if err != nil {
			return fmt.Errorf("Config error: %v", err)
		}

		if len(args) == 0 {
			return errors.New("Commit message file not provided")
		}

		path := args[0]

		message, err := commit.ReadCommitMessage(path)
		if err != nil {
			return fmt.Errorf("Error reading commit message: %v", err)
		}

		diff := ""
		if cfg.DiffLimit != nil {
			diff, err = git.GetStagedDiff()
			if err != nil {
				return fmt.Errorf("Error reading diff: %v", err)
			}
			diff = git.LimitDiff(diff, *cfg.DiffLimit)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*cfg.ImprovementRequestTimeout)*time.Second)
		defer cancel()

		provider, err := ai.NewProvider(ctx, cfg)
		if err != nil {
			return fmt.Errorf("Provider error: %v", err)
		}

		prompt := prompt.Build(message, diff, cfg.Language)

		spinner := ui.New("Improving commit message...\n")
		spinner.Start()
		defer spinner.Stop()

		improvedMessage, err := provider.ImproveCommitMessage(ctx, prompt)
		if err != nil {
			spinner.Stop()
			return fmt.Errorf("✖ Failed to improve commit: %v", err)
		}

		spinner.Stop()
		fmt.Print("✔ Commit message improved\n\n")

		ui.ShowPreview(message, improvedMessage)

		accepted, err := ui.Confirm("Apply improved commit message?", true)
		if err != nil {
			return err
		}

		if accepted {
			err = commit.WriteCommitMessage(path, improvedMessage)
			if err != nil {
				return err
			}
		}

		if !cfg.AllowFinalEdit {
			return nil
		}

		editCommitMessage, err := ui.Confirm("Do you want to make a final edit to the commit message?", true)
		if err != nil {
			return err
		}
		if editCommitMessage {
			err := editor.Open(path)
			if err != nil {
				return fmt.Errorf("Failed to open editor: %v", err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
