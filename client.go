package client

import (
	"github.com/GabeCordo/commandline"
	"github.com/GabeCordo/etl/client/controllers"
	"github.com/GabeCordo/etl/client/core"
)

func CommandLineClient() {
	// these are data files used by this executable to store metadata about created projects or ECDSA keys
	core.IfMissingInitializeFolders()

	profilePath := core.CliConfigFile()
	if commandLine := commandline.NewCommandLine(profilePath); commandLine != nil {

		// other commands
		commandLine.AddCommand("version", controllers.VersionCommand{})

		// cli core commands
		commandLine.AddCommand("key", controllers.KeyPairCommand{})
		commandLine.AddCommand("project", controllers.CreateProjectCommand{})
		commandLine.AddCommand("cluster", controllers.ClusterCommand{})
		commandLine.AddCommand("profile", controllers.ProfileCommand{})

		//local project interaction
		commandLine.AddCommand("deploy", controllers.DeployCommand{})
		commandLine.AddCommand("mount", controllers.MountCommand{})
		commandLine.AddCommand("endpoint", controllers.EndpointCommand{})
		commandLine.AddCommand("permission", controllers.PermissionCommand{})

		// remote project interaction

		// remote project observation (not complete)
		commandLine.AddCommand("interactive", controllers.InteractiveDashboardCommand{})

		commandLine.Run()
	}
}
