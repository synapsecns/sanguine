package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/contrib/docker-mirror/config"
	"github.com/synapsecns/sanguine/contrib/docker-mirror/mirror"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
)

var configFlag = &cli.StringFlag{
	Name:     "config",
	Usage:    "--config /path/to/config.yaml",
	Required: true,
}

var RunCommand = &cli.Command{
	Name:        "run",
	Description: "run the copier",
	Flags: []cli.Flag{
		configFlag,
	},
	Action: func(c *cli.Context) error {
		configLoc := core.ExpandOrReturnPath(c.String(configFlag.Name))
		configFile, err := os.ReadFile(configLoc)
		if err != nil {
			return fmt.Errorf("could not read config file: %w", err)
		}

		var cfg config.Config
		if err := yaml.Unmarshal(configFile, &cfg); err != nil {
			return fmt.Errorf("could not unmarshal config file: %w", err)
		}

		client, err := mirror.NewClient(&cfg)
		if err != nil {
			return fmt.Errorf("could not create client: %w", err)
		}

		_ = client

		return nil
	},
}
