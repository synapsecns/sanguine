package cmd

import (
	// used to embed markdown.
	_ "embed"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/urfave/cli/v2"
)

// Start starts the command line.
func Start(args []string) {
	app := cli.NewApp()
	app.Name = "scribe"
	app.Description = "scribe is used to run a generic event indexer"
	app.Usage = "scribe help"
	app.EnableBashCompletion = true

	// commands
	app.Commands = cli.Commands{infoCommand, backfillCommand}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
