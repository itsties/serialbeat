package main

import (
	"os"

	"github.com/itsties/serialbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
