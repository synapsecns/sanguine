package cmd

import (
	"flag"
	"fmt"

	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	cctpCmd "github.com/synapsecns/sanguine/services/cctp-relayer/cmd"
	"github.com/urfave/cli/v2"
)

// Start starts the command line tool.
func Start(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()
	app.Description = buildInfo.VersionString() + "Synapse RFQ Relayer Server"
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.EnableBashCompletion = true
	// TODO: should we really halt boot on because of metrics?
	app.Before = func(c *cli.Context) error {
		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}

	// check the embedded flag here to see if we should
	// include an embedded CCTP relayer command.
	flagSet := flag.NewFlagSet("RFQFlagSet", flag.ContinueOnError)
	embedded := flagSet.Bool(cctpCmd.EmbeddedFlag.Name, false, cctpCmd.EmbeddedFlag.DefaultText)
	err := flagSet.Parse(args)
	if err != nil {
		panic(fmt.Errorf("could not parse flags: %w", err))
	}

	// commands
	app.Commands = cli.Commands{runCommand}
	if *embedded {
		app.Commands = append(app.Commands, cctpCmd.RunCommand)
	}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err = app.Run(args)
	if err != nil {
		panic(err)
	}
}
