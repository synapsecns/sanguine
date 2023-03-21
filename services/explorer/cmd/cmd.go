package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/urfave/cli/v2"
)

// Start starts the command line.
func Start(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()
	app.Version = buildInfo.Version()

	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.Description = buildInfo.VersionString() + "An indexer + API serving platform analytics"
	app.Usage = "explorer help"
	app.EnableBashCompletion = true

	// TODO: should we really halt boot on because of metrics?
	app.Before = func(c *cli.Context) error {
		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}

	// commands
	app.Commands = cli.Commands{infoCommand, serverCommand, backfillCommand, livefillCommand}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
