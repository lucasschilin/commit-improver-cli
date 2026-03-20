package cmd

import (
	"errors"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			return errors.New("Not inside a git repository")
		}

		var (
			cfg *config.Config
			err error
		)

		switch {
		case globalFlag:
			cfg, err = config.LoadConfigFile(config.GlobalConfigPath())

		case repoUserFlag:
			cfg, err = config.LoadConfigFile(config.RepoUserConfigPath(repoRoot))

		case repoFlag:
			cfg, err = config.LoadConfigFile(config.RepoSharedConfigPath(repoRoot))

		default:
			cfg, err = config.Resolve(repoRoot)
		}

		if err != nil {
			return fmt.Errorf("Config error: %v", err)
		}

		output, err := config.ToYAML(cfg)
		if err != nil {
			return fmt.Errorf("Error serializing config: %v", err)
		}

		fmt.Println(output)

		return nil

	},
}

func init() {
	configCmd.AddCommand(configShowCmd)

	configShowCmd.Flags().BoolVar(&globalFlag, "global", false, "Show global config")
	configShowCmd.Flags().BoolVar(&repoFlag, "repo", false, "Show shared repo config")
	configShowCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Show user config for this repo")
}
