package cmd

import (
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/agents/notary/api"
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

// NotaryInfoCommand gets info about using the notary agent.
var NotaryInfoCommand = &cli.Command{
	Name:        "notary-info",
	Description: "learn how to use notary cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var metricsPortFlag = &cli.UintFlag{
	Name:  "metrics-port",
	Usage: "--port 5121",
	Value: 0,
}

var ignoreInitErrorsFlag = &cli.BoolFlag{
	Name:  "ignore-init-errors",
	Usage: "--ignore-init-errors",
	Value: false,
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/notary_config.yaml",
	TakesFile: true,
	Required:  true,
}

// NotaryRunCommand runs the notary.
var NotaryRunCommand = &cli.Command{
	Name:        "notary-run",
	Description: "runs the notary service",
	Flags:       []cli.Flag{configFlag, metricsPortFlag, ignoreInitErrorsFlag},
	Action: func(c *cli.Context) error {
		notaryConfig, err := config.DecodeNotaryConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		g, _ := errgroup.WithContext(c.Context)

		notary, err := notary.NewNotary(c.Context, notaryConfig)
		if err != nil && !c.Bool(ignoreInitErrorsFlag.Name) {
			return fmt.Errorf("failed to create notary: %w", err)
		}

		g.Go(func() error {
			err = notary.Start(c.Context)
			if err != nil && !c.Bool(ignoreInitErrorsFlag.Name) {
				return fmt.Errorf("failed to run notary: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			err := api.Start(c.Context, uint16(c.Uint(metricsPortFlag.Name)))
			if err != nil {
				return fmt.Errorf("failed to start api: %w", err)
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return fmt.Errorf("failed to run notary: %w", err)
		}

		return nil
	},
}

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}
