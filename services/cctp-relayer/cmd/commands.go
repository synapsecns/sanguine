package cmd

import "github.com/urfave/cli/v2"

// runCommand runs the cctp relayer
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the cctp relayer",
	Action: func(c *cli.Context) error {
		// TODO
		return nil
	},
}
