package controllers

import (
	"fmt"
	"github.com/GabeCordo/commandline"
	etlCore "github.com/GabeCordo/etl/core"
	clientCore "github.com/GabeCordo/etlclient/core"
	"github.com/GabeCordo/fack"
	"github.com/GabeCordo/toolchain/files"
)

type EndpointCommand struct {
}

func (ec EndpointCommand) Run(cl *commandline.CommandLine) commandline.TerminateOnCompletion {
	if cl.Flag(commandline.Create) {
		ec.CreateKeyEndpoint(cl)
	} else if cl.Flag(commandline.Show) {
		ec.ShowEndpoints(cl)
	}

	return true // done
}

func (ec EndpointCommand) CreateKeyEndpoint(cl *commandline.CommandLine) commandline.TerminateOnCompletion {
	keyIdentity := cl.NextArg()
	if keyIdentity == commandline.FinalArg {
		fmt.Println("missing key identifier")
		return true
	}

	configPath := files.EmptyPath().Dir(cl.MetaData.WorkingDirectory).File("config.etl.json")
	if configPath.DoesNotExist() {
		fmt.Println("no etl project found")
		return true
	}

	keysFilePath := clientCore.EtlKeysFile()
	if keysFilePath.DoesNotExist() {
		fmt.Println("missing key metadata folder, try restarting")
		return true
	}

	keysFile := clientCore.JSONToKeysFile(keysFilePath)
	if _, found := keysFile.Keys[keyIdentity]; !found {
		fmt.Println("no key exists with the identifier " + keyIdentity)
		return true
	}

	config := etlCore.NewConfig("tmp")
	etlCore.JSONToETLConfig(config, configPath.ToString())

	if _, found := config.Auth.Trusted[keyIdentity]; found {
		fmt.Println("a trusted already exists with this identifier")
		return true
	}

	endpoint := fack.NewEndpoint(keyIdentity, nil)
	endpoint.X509 = keysFile.Keys[keyIdentity].PublicKey
	if config.Auth.Trusted == nil {
		config.Auth.Trusted = make(map[string]*fack.Endpoint) // Todo - this is very bad practice, please follow SRP
	}
	config.Auth.Trusted[keyIdentity] = endpoint

	config.ToJson(configPath.ToString())

	return true // done
}

func (ec EndpointCommand) ShowEndpoints(cl *commandline.CommandLine) commandline.TerminateOnCompletion {
	configPath := files.EmptyPath().Dir(cl.MetaData.WorkingDirectory).File("config.etl.json")
	if configPath.DoesNotExist() {
		fmt.Println("no etl project found")
		return true
	}

	config := etlCore.NewConfig("tmp")
	etlCore.JSONToETLConfig(config, configPath.ToString())

	if config.Auth.Trusted == nil {
		fmt.Println("no trusted endpoints")
		return true
	}

	for identity, endpoint := range config.Auth.Trusted {
		fmt.Printf("id: %s\n", identity)
		fmt.Printf("\tglobal-permissions: %s\n", endpoint.GlobalPermissions.String())
		if (endpoint.LocalPermissions != nil) && (len(endpoint.LocalPermissions) > 0) {
			fmt.Println("\tlocal-permissions:")
			for route, permission := range endpoint.LocalPermissions {
				fmt.Printf("\t\t/%s => %s\n", route, permission.String())
			}
		}
		fmt.Println()
	}

	return true // done
}
