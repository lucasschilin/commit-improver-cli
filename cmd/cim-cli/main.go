package main

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
