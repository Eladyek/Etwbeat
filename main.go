package main

import (
	"os"

	"github.com/eladyek/Etwbeat/cmd"

	_ "github.com/eladyek/Etwbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
