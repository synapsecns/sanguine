package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/committee/devnet/provisioner"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the committee devnet",
	Flags: []cli.Flag{
		configFlag,
		&commandline.LogLevel,
	},
	Action: func(c *cli.Context) error {
		commandline.SetLogLevel(c)

		input, err := os.ReadFile(filepath.Clean(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		var cfg config.Config
		// TODO: consider moving this for marshal/unmarshall tests
		err = yaml.Unmarshal(input, &cfg)
		if err != nil {
			return fmt.Errorf("could not unmarshal config: %w", err)
		}

		metricsProvider := metrics.Get()

		provisioner, err := provisioner.NewProvisioner(c.Context, metricsProvider, cfg)
		if err != nil {
			return fmt.Errorf("could not create provisioner: %w", err)
		}

		return provisioner.Run(c.Context, cfg)
	},
}
