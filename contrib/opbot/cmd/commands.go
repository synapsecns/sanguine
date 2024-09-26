package cmd

import (
	"fmt"
	"os"

	"github.com/synapsecns/sanguine/contrib/opbot/botmd"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"

	// used for testing.

	_ "github.com/joho/godotenv/autoload"
	"github.com/synapsecns/sanguine/contrib/opbot/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var fileFlag = &cli.StringFlag{
	Name:  "config",
	Usage: "--config /Users/synapsecns/config.yaml",
}

var slackBotCommand = &cli.Command{
	Name:  "start",
	Usage: "start the slack bot",
	Flags: []cli.Flag{fileFlag},
	Action: func(c *cli.Context) error {
		configFile, err := os.ReadFile(core.ExpandOrReturnPath(c.String(fileFlag.Name)))
		if err != nil {
			return fmt.Errorf("failed to open config file: %w", err)
		}

		var cfg config.Config
		err = yaml.Unmarshal(configFile, &cfg)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config file: %w", err)
		}

		if cfg.SlackAppToken == "" {
			return fmt.Errorf("slack app token is required")
		}

		if cfg.SlackBotToken == "" {
			return fmt.Errorf("slack bot token is required")
		}

		botServer := botmd.NewBot(metrics.Get(), cfg)
		err = botServer.Start(c.Context)
		if err != nil {
			return fmt.Errorf("failed to start bot: %w", err)
		}

		return nil
	},
}
