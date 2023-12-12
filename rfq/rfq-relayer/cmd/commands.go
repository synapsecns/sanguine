package cmd

import (
	"fmt"
	"os"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/sql/mysql"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service" // used to embed markdown.

	"github.com/urfave/cli/v2"
)

// infoCommand gets info about using the scribe service.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use rfq relayer cli",
	Action: func(c *cli.Context) error {
		fmt.Println("Run relayer --config path/to/config.yaml to get sterted") // TODO: more info here w/markdown
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /path/to/config.yaml",
	TakesFile: true,
	Required:  true,
}
var relayerCommand = &cli.Command{
	Name:        "relayer",
	Description: "starts the relayer on all configured chains",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		relayerConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)
		}
		// Create MySQL Database connection
		metricHandler := metrics.Get()
		dbname := os.Getenv("MYSQL_DATABASE")
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
		mysqlStore, err := mysql.NewMysqlStore(c.Context, connString, metricHandler, relayerConfig.SkipMigrations)
		if err != nil {
			return fmt.Errorf("failed to create mysql store: %w", err)
		}
		metricHandler.AddGormCallbacks(mysqlStore.DB())

		relayer, err := service.NewRelayer(c.Context, &relayerConfig, mysqlStore, metricHandler)
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
