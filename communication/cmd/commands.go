package cmd

import (
	"github.com/synapsecns/sanguine/communication/config"
	"github.com/synapsecns/sanguine/communication/node"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/core/metrics"
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
			return err
		}

		metricsProvider := metrics.Get()

		var cfg config.Config
		err = yaml.Unmarshal(input, &cfg)
		if err != nil {
			return err
		}

		createdNode, err := node.NewNode(c.Context, metricsProvider, cfg)
		if err != nil {
			return err
		}

		err = createdNode.Start(c.Context)
		if err != nil {
			return err
		}
		return nil
	},
}
