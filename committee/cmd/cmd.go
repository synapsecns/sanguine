// Package cmd provides the command line interface for the committee node.
package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/urfave/cli/v2"
)

// Start starts the committee node.
func Start(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()
	app.Description = buildInfo.VersionString() + "committee provider for synapse"
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.EnableBashCompletion = true

	app.Before = func(c *cli.Context) error {
		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}

	// commands
	app.Commands = cli.Commands{runCommand}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
