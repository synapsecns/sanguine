package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	explorerConfig "github.com/synapsecns/sanguine/services/sin-explorer/config"
	explorer "github.com/synapsecns/sanguine/services/sin-explorer/indexer"
	"github.com/synapsecns/sanguine/services/sin-explorer/metadata"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"os"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

var indexerCommand = &cli.Command{
	Name:        "indexer",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		configContents, err := os.ReadFile(c.String(configFlag.Name))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		var config explorerConfig.Config
		err = yaml.Unmarshal(configContents, &config)
		if err != nil {
			return fmt.Errorf("could not unmarshal config file: %w", err)
		}

		m, err := metrics.NewFromEnv(c.Context, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("could not create metrics: %w", err)
		}
		indexer, err := explorer.NewIndexer(c.Context, config, m)
		if err != nil {
			return fmt.Errorf("could not create indexer: %w", err)
		}

		err = indexer.Start(c.Context)
		if err != nil {
			return fmt.Errorf("could not start indexer: %w", err)
		}

		return nil
	},
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {

		return nil
	},
}
