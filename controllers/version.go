package controllers

import (
	"fmt"
	"github.com/GabeCordo/commandline"
	"github.com/GabeCordo/etl/client/core"
)

// VERSION COMMAND START

type VersionCommand struct {
}

func (vc VersionCommand) Run(cl *commandline.CommandLine) commandline.TerminateOnCompletion {
	fmt.Println(core.Version(cl))
	return true
}
