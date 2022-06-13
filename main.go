package main

import (
	"os"

	"github.com/davdjl/timebeat/cmd"

	_ "github.com/davdjl/timebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
