package main

import (
	"os"

	"github.com/lucasschilin/cim-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
