package cmd

import (
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/urfave/cli/v2"
)

const appName = "omnirpc"

// Start starts the command line.
func Start(args []string) {
	app := cli.NewApp()
	app.Name = appName
	app.Version = config.AppVersion
	app.Description = "Used for checking the lowest latency rpc endpoint fora given chain"
	app.Commands = []*cli.Command{latencyCommand, chainListCommand, publicConfigCommand, serverCommand, debugResponse}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
