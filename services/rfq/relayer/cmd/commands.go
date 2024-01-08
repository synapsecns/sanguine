// Package cmd provides the command line interface for the RFQ relayer service
package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

// runCommand runs the cctp relayer.
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the API Server",
	Flags:       []cli.Flag{configFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		commandline.SetLogLevel(c)
		cfg, err := relconfig.LoadConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		metricsProvider := metrics.Get()

		relayer, err := service.NewRelayer(c.Context, metricsProvider, cfg)
		if err != nil {
			return fmt.Errorf("could not create relayer: %w", err)
		}

		err = relayer.Start(c.Context)
		if err != nil {
			return fmt.Errorf("could not start relayer: %w", err)
		}
		return nil
	},
}
