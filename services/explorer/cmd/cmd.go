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
	app.Name = "explorer"
	app.Description = "An indexer + API serving platform analytics"
	app.Usage = "explorer help"
	app.EnableBashCompletion = true

	// commands
	app.Commands = cli.Commands{infoCommand, placeholderCommand}


	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
