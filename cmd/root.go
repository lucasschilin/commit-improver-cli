package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cim-cli",
	Short: "Commit Improver CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Commit Improver CLI running")
	},
}

func Execute() error {

	return rootCmd.Execute()

}
