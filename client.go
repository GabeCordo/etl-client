package main

import (
	"github.com/GabeCordo/commandline"
	"github.com/GabeCordo/etlclient/controllers"
	"github.com/GabeCordo/etlclient/core"
)

func CommandLineClient() {
	// these are data files used by this executable to store metadata about created projects or ECDSA keys
	core.IfMissingInitializeFolders()

	profilePath := core.CliConfigFile()
	if commandLine := commandline.NewCommandLine(profilePath); commandLine != nil {

		// other commands
		commandLine.AddCommand("version", controllers.VersionCommand{}).SetCategory("core")

		// cli core commands
		commandLine.AddCommand("key", controllers.KeyPairCommand{}).SetCategory("etl")
		commandLine.AddCommand("project", controllers.CreateProjectCommand{}).SetCategory("etl")
		commandLine.AddCommand("cluster", controllers.ClusterCommand{}).SetCategory("etl")
		commandLine.AddCommand("profile", controllers.ProfileCommand{}).SetCategory("etl")

		//local project interaction
		commandLine.AddCommand("deploy", controllers.DeployCommand{}).SetCategory("local project")
		commandLine.AddCommand("mount", controllers.MountCommand{}).SetCategory("local project")
		commandLine.AddCommand("endpoint", controllers.EndpointCommand{}).SetCategory("local project")
		commandLine.AddCommand("permission", controllers.PermissionCommand{}).SetCategory("local project")

		// remote project interaction

		// remote project observation (not complete)
		commandLine.AddCommand("dashboard", controllers.InteractiveDashboardCommand{}).SetCategory("visuals")

		commandLine.Run()
	}
}
