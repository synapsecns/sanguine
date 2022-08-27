package cmd

import (
	"github.com/gen2brain/beeep"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/synapse-node/config"
	"github.com/urfave/cli/v2"
)

// appName is the name of the abi generator.
const appName = "abigen"

// Run runs the cmd
// TODO maybe use ifacemaker to generate interfaces for these.
func Run(args []string) {
	app := cli.NewApp()
	app.Name = appName
	app.Description = "abi generator. This extends the standard abi gen by requiring a sol-version and using docker. It also generates metadata for each contract."
	app.Commands = []*cli.Command{
		GenerateCommand,
		EtherscanCommand,
	}

	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		// we send an additional alert through beep because go:generate *will* silently fail if ran as
		// go:generate ./...
		logoPath, _ := config.GetLogoPath()

		_ = beeep.Notify("AbiGen Failed", err.Error(), logoPath)
		panic(err)
	}
}
