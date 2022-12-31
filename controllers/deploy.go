package controllers

import (
	"github.com/GabeCordo/commandline"
	"github.com/GabeCordo/etl/core"
	"github.com/GabeCordo/toolchain/files"
	"log"
	"os"
	"os/exec"
)

// DEPLOY COMMAND START

type DeployCommand struct {
}

func (dc DeployCommand) Run(cl *commandline.CommandLine) commandline.TerminateOnCompletion {

	if cl.Flag(commandline.Debug) {
		log.Println("(+) starting up etl")
	}

	projectConfigFilePath := files.EmptyPath().Dir(cl.MetaData.WorkingDirectory).File("config.etl.json")
	if projectConfigFilePath.DoesNotExist() {
		log.Println("you are not in an etl project")
		return true
	}

	projectConfig := core.NewConfig("temp")
	core.JSONToETLConfig(projectConfig, projectConfigFilePath.ToString())

	entryPointFile := projectConfig.Name + ".root.go"
	mainPath := files.EmptyPath().File(entryPointFile)
	if _, err := os.Stat(mainPath.ToString()); err == nil {
		// if the file exists run the main module
		runEtlMainCmd := exec.Command("go run " + entryPointFile)
		if err = runEtlMainCmd.Run(); err != nil {
			// there was an source error inside the etl project
			log.Print(err)
		}
	} else {
		// if the file does not exists, let them know that they are not in an etl project folder
		log.Println("(!) you are not in an ETL project")
	}

	if cl.Flag(commandline.Debug) {
		log.Println("(-) shutting down etl")
	}

	return true // end of program
}
