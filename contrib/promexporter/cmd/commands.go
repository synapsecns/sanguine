package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/contrib/promexporter/config"
	"github.com/synapsecns/sanguine/contrib/promexporter/exporters"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

// TODO: add a command to print default config.
var exporterCommand = &cli.Command{
	Name:  "start",
	Usage: "start the prometheus exporter",
	Flags: []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		exporterConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)
		}

		fmt.Println("starting server")

		err = exporters.StartExporterServer(c.Context, metrics.Get(), exporterConfig)
		if err != nil {
			return fmt.Errorf("could not start exporter server: %w", err)
		}
		return nil
	},
}
