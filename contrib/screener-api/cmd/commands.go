package cmd

import (
	"fmt"
	"os"

	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var fileFlag = &cli.StringFlag{
	Name:  "config",
	Usage: "--config /Users/synapsecns/config.yaml",
}

var screenerCommand = &cli.Command{
	Name:  "screener",
	Usage: "screener",
	Flags: []cli.Flag{fileFlag},
	Action: func(c *cli.Context) error {
		configFile, err := os.ReadFile(c.String(fileFlag.Name))
		if err != nil {
			return fmt.Errorf("failed to open config file: %w", err)
		}

		var cfg config.Config
		err = yaml.Unmarshal(configFile, &cfg)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config file: %w", err)
		}

		screnr, err := screener.NewScreener(c.Context, cfg, metrics.Get())
		if err != nil {
			return fmt.Errorf("failed to create screener: %w", err)
		}

		err = screnr.Start(c.Context)
		if err != nil {
			return fmt.Errorf("failed to start screener: %w", err)
		}

		return nil
	},
}

var inFileFlag = &cli.StringFlag{
	Name:  "in-file",
	Usage: "Specify the path to the input CSV file. Example: --in-file /path/to/in.csv",
}

var outDirFlag = &cli.StringFlag{
	Name:  "out-dir",
	Usage: "Specify the path to the output directory where split CSV files will be saved. Example: --out-dir /path/to/output",
}
