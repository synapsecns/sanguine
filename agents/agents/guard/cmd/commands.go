package cmd

import (
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/agents/agents/guard"
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

// GuardInfoCommand gets info about using the guard agent.
var GuardInfoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use guard cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/guard_config.yaml",
	TakesFile: true,
	Required:  true,
}

// GuardRunCommand runs the guard.
var GuardRunCommand = &cli.Command{
	Name:        "run",
	Description: "runs the guard service",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		guardConfig, err := config.DecodeGuardConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		g, _ := errgroup.WithContext(c.Context)

		guard, err := guard.NewGuard(c.Context, guardConfig)
		if err != nil {
			return fmt.Errorf("failed to create guard: %w", err)
		}

		g.Go(func() error {
			err = guard.Start(c.Context)
			if err != nil {
				return fmt.Errorf("failed to run guard: %w", err)
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return fmt.Errorf("failed to run guard: %w", err)
		}

		return nil
	},
}
