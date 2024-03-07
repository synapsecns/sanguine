package cmd

import (
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
	fmt.Printf("Raw args: %v\n", args)
	app.Before = func(c *cli.Context) error {
		fmt.Printf("Running 'before' setup with flags: %v\n", c.Command.Flags)
		if c.Bool(cctpCmd.EmbeddedFlag.Name) {
			fmt.Println("Running as embedded service")
			app.Commands = append(app.Commands, cctpCmd.RunCommand)
		} else {
			fmt.Println("Not running as embedded service")
		}

		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}

	// commands
	app.Commands = cli.Commands{runCommand}
	fmt.Println("Created commands")
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
