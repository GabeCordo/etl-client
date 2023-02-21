package core

import (
	"fmt"
	"github.com/GabeCordo/commandline"
	"github.com/GabeCordo/toolchain/files"
	"os"
	"strings"
	"time"
)

func Version(commandLine *commandline.CommandLine) string {
	strVersion := fmt.Sprintf("%.2f", commandLine.Config.Version)
	strTimeNow := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
	return "ETLFramework Version " + strVersion + " " + strTimeNow
}

func RootEtlFolder() files.Path {
	executableFilePath, _ := os.Executable()
	t := strings.Split(executableFilePath, files.PathSeparator())
	executableFilePath = strings.Join(t[:len(t)-1], files.PathSeparator())
	return files.EmptyPath().Dir(executableFilePath) // remove "/build/etl" from the end of the path
}

func TemplateFolder() files.Path {
	return RootEtlFolder().Dir(".templates")
}

func DataFolder() files.Path {
	return RootEtlFolder().Dir(".data")
}

func CliConfigFile() files.Path {
	return DataFolder().File("config.cli.json")
}

func EtlClientFile() files.Path {
	return DataFolder().File("client.etl.json")
}

func EtlKeysFile() files.Path {
	return DataFolder().File("keys.etl.json")
}

func IfMissingInitializeFolders() {

	dataFolderPath := DataFolder()
	if dataFolderPath.DoesNotExist() {
		if err := dataFolderPath.MkDir(); err != nil {
			fmt.Println(err)
		}
	}

	cliConfigFilePath := CliConfigFile()
	if cliConfigFilePath.DoesNotExist() {
		cliConfigFilePath.Create()

		fmt.Println(cliConfigFilePath.ToString())

		config := commandline.NewConfig()
		config.ToJson(cliConfigFilePath)
	}

	projectsFilePath := EtlClientFile()
	if projectsFilePath.DoesNotExist() {
		projectsFilePath.Create()

		config := NewConfig()
		config.ToJson(projectsFilePath)
	}

	keysFilePath := EtlKeysFile()
	if keysFilePath.DoesNotExist() {
		os.Create(keysFilePath.ToString())

		config := NewKeysFile()
		config.ToJson(keysFilePath)
	}
}
