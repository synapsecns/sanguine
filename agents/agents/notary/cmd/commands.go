package cmd

import (
	"sync/atomic"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/hedzr/log"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/agents/notary/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"golang.org/x/sync/errgroup"

	// used to embed markdown.
	_ "embed"
	"fmt"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/core"
	scribeAPI "github.com/synapsecns/sanguine/services/scribe/api"
	scribeCmd "github.com/synapsecns/sanguine/services/scribe/cmd"
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
		notaryConfig, err := config.DecodeAgentConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to decode config: %w", err)
		}

		var shouldRetryAtomic atomic.Bool
		shouldRetryAtomic.Store(true)

		for shouldRetryAtomic.Load() {
			shouldRetryAtomic.Store(false)

			var scribeClient client.ScribeClient

			g, _ := errgroup.WithContext(c.Context)

			eventDB, err := scribeAPI.InitDB(c.Context, "mysql", "root:MysqlPassword@tcp(agents-mysql:3306)/scribe?parseTime=true")
			if err != nil {
				return fmt.Errorf("failed to initialize database: %w", err)
			}

			scribeClients := make(map[uint32][]backfill.ScribeBackend)

			for _, domain := range notaryConfig.Domains {
				for confNum := 1; confNum <= scribeCmd.MaxConfirmations; confNum++ {
					chainID := domain.DomainID
					backendClient, err := backfill.DialBackend(c.Context, fmt.Sprintf("%s/%d/rpc/%d", "https://rpc.interoperability.institute/confirmations", confNum, chainID))
					if err != nil {
						return fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/1/rpc/%d", "https://rpc.interoperability.institute/confirmations", chainID))
					}

					scribeClients[chainID] = append(scribeClients[chainID], backendClient)
				}
			}

			scribe, err := node.NewScribe(eventDB, scribeClients, notaryConfig.EmbeddedScribeConfig)
			if err != nil {
				return fmt.Errorf("failed to initialize scribe: %w", err)
			}

			g.Go(func() error {
				err := scribe.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start scribe: %w", err)
				}

				return nil
			})

			embedded := client.NewEmbeddedScribe("mysql", "root:MysqlPassword@tcp(agents-mysql:3306)/scribe?parseTime=true")

			g.Go(func() error {
				err := embedded.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start embedded scribe: %w", err)
				}

				return nil
			})

			scribeClient = embedded.ScribeClient

			notaryConfig.ScribeURL = scribeClient.URL
			notaryConfig.ScribePort = uint32(scribeClient.Port)
			notary, err := notary.NewNotary(c.Context, notaryConfig)
			if err != nil && !c.Bool(ignoreInitErrorsFlag.Name) {
				return fmt.Errorf("failed to create notary: %w", err)
			}

			g.Go(func() error {
				err = notary.Start(c.Context)
				if err != nil {
					shouldRetryAtomic.Store(true)

					log.Errorf("Error running guard, will sleep for a minute and retry: %v", err)
					time.Sleep(60 * time.Second)
					return fmt.Errorf("failed to create notary: %w", err)
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
		}

		return nil
	},
}

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}
