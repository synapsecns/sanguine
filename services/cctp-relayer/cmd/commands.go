package cmd

import (
	"fmt"

	"github.com/synapsecns/sanguine/core/commandline"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

var dbFlag = &cli.StringFlag{
	Name:     "db",
	Usage:    "--db <sqlite> or <mysql>",
	Value:    "sqlite",
	Required: true,
}

var pathFlag = &cli.StringFlag{
	Name:     "path",
	Usage:    "--path <path/to/database> or <database url>",
	Required: true,
}

var scribePortFlag = &cli.UintFlag{
	Name:  "scribe-port",
	Usage: "--scribe-port <port>",
	Value: 0,
}

var scribeURL = &cli.StringFlag{
	Name:  "scribe-url",
	Usage: "--scribe-url <url>",
}

// runCommand runs the cctp relayer.
var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the cctp relayer",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag, scribePortFlag, scribeURL, &commandline.LogLevel},
	Action: func(c *cli.Context) (err error) {
		commandline.SetLogLevel(c)
		cfg, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		_, err = cfg.IsValid(c.Context)
		if err != nil {
			return fmt.Errorf("could not decode config file: %w", err)
		}

		dbTypeFromString, err := dbcommon.DBTypeFromString(c.String(dbFlag.Name))
		if err != nil {
			return fmt.Errorf("could not get db type from string: %w", err)
		}

		path := core.ExpandOrReturnPath(c.String(pathFlag.Name))

		metricsProvider := metrics.Get()

		store, err := sql.Connect(c.Context, dbTypeFromString, path, metricsProvider)
		if err != nil {
			return fmt.Errorf("could not connect to database: %w", err)
		}

		scribeClient := client.NewRemoteScribe(uint16(c.Uint(scribePortFlag.Name)), c.String(scribeURL.Name), metricsProvider).ScribeClient
		omnirpcClient := omniClient.NewOmnirpcClient(cfg.BaseOmnirpcURL, metricsProvider, omniClient.WithCaptureReqRes())
		attAPI := attestation.NewCircleAPI(c.String(cfg.CircleAPIURl))

		cctpRelayer, err := relayer.NewCCTPRelayer(c.Context, cfg, store, scribeClient, omnirpcClient, metricsProvider, attAPI)
		if err != nil {
			return fmt.Errorf("could not create cctp relayer: %w", err)
		}

		err = cctpRelayer.Run(c.Context)
		if err != nil {
			return fmt.Errorf("could not run cctp relayer: %w", err)
		}
		return nil
	},
}
