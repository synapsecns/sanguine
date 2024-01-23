// Package cmd provides the command line interface for the RFQ API service.
package cmd

import (
	"fmt"

	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/dbcommon"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/stiprelayer/db/sql"
	"github.com/synapsecns/sanguine/services/stiprelayer/relayer"
	"github.com/synapsecns/sanguine/services/stiprelayer/stipconfig"
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
		cfg, err := stipconfig.LoadConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		metricsProvider := metrics.Get()

		dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
		if err != nil {
			return fmt.Errorf("could not get db type: %w", err)
		}
		store, err := sql.Connect(c.Context, dbType, cfg.Database.DSN, metricsProvider)
		if err != nil {
			return fmt.Errorf("could not connect to database: %w", err)
		}

		omnirpcClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricsProvider, omniClient.WithCaptureReqRes())
		stipRelayer, err := relayer.NewSTIPRelayer(c.Context, cfg, metricsProvider, omnirpcClient, store)
		if err != nil {
			return fmt.Errorf("could not create api server: %w", err)
		}

		err = stipRelayer.Run(c.Context)
		if err != nil {
			return fmt.Errorf("could not run cctp relayer: %w", err)
		}
		return nil
	},
}
