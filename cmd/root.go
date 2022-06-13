package cmd

import (
	"github.com/davdjl/timebeat/beater"

	cmd "github.com/elastic/beats/v7/libbeat/cmd"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
)

// Name of this beat
var Name = "timebeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
