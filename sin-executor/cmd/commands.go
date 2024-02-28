package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/sin-executor/config"
	"github.com/synapsecns/sanguine/sin-executor/executor"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "path to the config file",
	TakesFile: true,
}

var runCommand = &cli.Command{
	Name:        "run",
	Description: "run the committee node",
	Flags: []cli.Flag{
		configFlag,
		&commandline.LogLevel,
	},
	Action: func(c *cli.Context) error {
		commandline.SetLogLevel(c)

		input, err := os.ReadFile(filepath.Clean(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}

		metricsProvider := metrics.Get()

		var cfg config.Config
		err = yaml.Unmarshal(input, &cfg)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config: %w", err)
		}

		exec, err := executor.NewExecutor(c.Context, metricsProvider, cfg)
		if err != nil {
			return fmt.Errorf("failed to create executor: %w", err)
		}

		err = exec.Start(c.Context)
		if err != nil {
			return fmt.Errorf("failed to start executor: %w", err)
		}

		return nil
	},
}
