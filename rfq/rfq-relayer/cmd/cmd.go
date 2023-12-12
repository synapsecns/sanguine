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
	app.Description = buildInfo.VersionString() + " - RFQ Relayer"
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.EnableBashCompletion = true
	app.Before = func(c *cli.Context) error {
		err := metrics.Setup(c.Context, buildInfo)
		if err != nil {
			return fmt.Errorf("failed to setup metrics: %w", err)
		}
		return nil
	}

	// commands
	app.Commands = cli.Commands{infoCommand, relayerCommand}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	// Run
	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
