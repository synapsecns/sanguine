package cmd

import (
	"fmt"
	"strings"

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

	// TODO: there should be a cleaner way to parse the 'embedded' flag outside of run command
	embedded := false
	for _, arg := range args {
		if strings.Contains(arg, cctpCmd.EmbeddedFlag.Name) {
			fmt.Printf("Found 'embedded' arg: %v; running CCTP relayer as embedded service\n", arg)
			embedded = true
			break
		}
	}

	// TODO: should we really halt boot on because of metrics?
	app.Before = func(c *cli.Context) error {
		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}

	// commands
	app.Commands = cli.Commands{runCommand}
	if embedded {
		app.Commands = append(app.Commands, cctpCmd.RunCommand)
	}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
