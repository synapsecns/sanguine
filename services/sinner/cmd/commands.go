package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	"golang.org/x/sync/errgroup"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/api"
	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	unifiedConfig "github.com/synapsecns/sanguine/services/sinner/config/unified"

	"github.com/synapsecns/sanguine/services/sinner/service"

	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// infoCommand references the help info from the cmd.md file and presents it.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use sinner cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		fmt.Println("port", c.Uint("port"))
		decodeConfig, err := serverConfig.DecodeServerConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)
		}

		err = api.Start(c.Context, decodeConfig, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}

		return nil
	},
}

// nolint:dupl
var indexerCommand = &cli.Command{
	Name:        "indexer",
	Description: "indexes contracts from config",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := indexerConfig.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}
		db, err := api.InitDB(c.Context, decodeConfig.DBType, decodeConfig.DBPath, metrics.Get(), decodeConfig.SkipMigrations)
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}

		sinnerService, err := service.NewSinner(db, decodeConfig, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not create Sinner indexer: %w", err)
		}
		err = sinnerService.Index(c.Context)
		if err != nil {
			return fmt.Errorf("could not backfill backfiller: %w", err)
		}
		return nil
	},
}

// nolint:dupl
var unifiedCommand = &cli.Command{
	Name:        "unified",
	Description: "runs both the Sinner server and indexer",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := unifiedConfig.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}
		db, err := api.InitDB(c.Context, decodeConfig.DBType, decodeConfig.DBPath, metrics.Get(), decodeConfig.SkipMigrations)
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}

		g, groupCtx := errgroup.WithContext(c.Context)
		// Start indexer
		g.Go(func() error {
			sinnerService, err := service.NewSinner(db, decodeConfig.IndexerConfig(), metrics.Get())
			if err != nil {
				return fmt.Errorf("could not create Sinner indexer: %w", err)
			}
			err = sinnerService.Index(groupCtx)
			if err != nil {
				return fmt.Errorf("could not backfill backfiller: %w", err)
			}
			return nil
		})
		// Start server
		g.Go(func() error {
			err = api.Start(groupCtx, decodeConfig.ServerConfig(), metrics.Get())
			if err != nil {
				return fmt.Errorf("could not start server: %w", err)
			}
			return nil
		})

		err = g.Wait()
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	},
}
