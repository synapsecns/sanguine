// Package cmd provides the command line interface for the RFQ guard service
package cmd

import (
	"fmt"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/guard/guardconfig"
	"github.com/synapsecns/sanguine/services/rfq/guard/service"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

// runCommand runs the rfq guard.
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the guard",
	Flags:       []cli.Flag{configFlag, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		commandline.SetLogLevel(c)
		cfg, err := guardconfig.LoadConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		metricsProvider := metrics.Get()

		guard, err := service.NewGuard(c.Context, metricsProvider, cfg, nil)
		if err != nil {
			return fmt.Errorf("could not create guard: %w", err)
		}

		err = guard.Start(c.Context)
		if err != nil {
			return fmt.Errorf("could not start guard: %w", err)
		}
		return nil
	},
}
