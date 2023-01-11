package cmd

import (
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"golang.org/x/sync/errgroup"

	// used to embed markdown.
	_ "embed"
	"fmt"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// infoCommand gets info about using the notary agent.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use notary cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/notary_config.yaml",
	TakesFile: true,
	Required:  true,
}

var dbFlag = &cli.StringFlag{
	Name:     "db",
	Usage:    "--db <sqlite> or <mysql>",
	Value:    "sqlite",
	Required: true,
}

var pathFlag = &cli.StringFlag{
	Name:     "path",
	Usage:    "--path <path/to/database> or <database url>",
	Value:    "",
	Required: true,
}

var runCommand = &cli.Command{
	Name:        "run",
	Description: "runs the executor service",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		notaryConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		g, _ := errgroup.WithContext(c.Context)

		notary, err := notary.NewNotary(c.Context, notaryConfig)
		if err != nil {
			return fmt.Errorf("failed to create notary: %w", err)
		}

		g.Go(func() error {
			err = notary.Start(c.Context)
			if err != nil {
				return fmt.Errorf("failed to run notary: %w", err)
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return fmt.Errorf("failed to run notary: %w", err)
		}

		return nil
	},
}
