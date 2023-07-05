package cmd

import (
	"github.com/synapsecns/sanguine/agents/agents/guard/metadata"
	"github.com/synapsecns/sanguine/core/metrics"
	"sync/atomic"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/hedzr/log"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/guard/api"
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
	Name:        "guard-info",
	Description: "learn how to use guard cli",
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

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/guard_config.yaml",
	TakesFile: true,
	Required:  true,
}

// GuardRunCommand runs the guard.
var GuardRunCommand = &cli.Command{
	Name:        "guard-run",
	Description: "runs the guard service",
	Flags:       []cli.Flag{configFlag, metricsPortFlag},
	Action: func(c *cli.Context) error {
		handler, err := metrics.NewFromEnv(c.Context, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("failed to create metrics handler: %w", err)
		}
		guardConfig, err := config.DecodeAgentConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		var shouldRetryAtomic atomic.Bool
		shouldRetryAtomic.Store(true)

		for shouldRetryAtomic.Load() {
			shouldRetryAtomic.Store(false)

			g, _ := errgroup.WithContext(c.Context)

			guard, err := guard.NewGuard(c.Context, guardConfig, handler)
			if err != nil {
				return fmt.Errorf("failed to create guard: %w", err)
			}

			g.Go(func() error {
				err = guard.Start(c.Context)
				if err != nil {
					shouldRetryAtomic.Store(true)

					log.Errorf("Error running guard, will sleep for a minute and retry: %v", err)
					time.Sleep(60 * time.Second)
					return fmt.Errorf("failed to run guard: %w", err)
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
				return fmt.Errorf("failed to run guard: %w", err)
			}
		}

		return nil
	},
}

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}
