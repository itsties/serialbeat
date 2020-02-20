package cmd

import (
	"github.com/itsties/serialbeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
        "github.com/elastic/beats/libbeat/cmd/instance"

)

// Name of this beat
var Name = "serialbeat"

// RootCmd to handle beats cli
//var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
