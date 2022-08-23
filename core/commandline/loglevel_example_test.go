package commandline_test

import (
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger = log.Logger("synapse-logger-example")

func ExampleSetLogLevel() {
	app := cli.NewApp()
	app.Commands = cli.Commands{
		{
			Name:        "test",
			Description: "I'm a test command",
			Flags: []cli.Flag{
				&commandline.LogLevel,
			},
			Action: func(c *cli.Context) error {
				commandline.SetLogLevel(c)
				logger.Debug("I won't be shown if level is set to warn")
				logger.Warn("I will be shown if level is set to warn")
				return nil
			},
		},
	}
	fmt.Printf("running ./example %s --%s %s", app.Commands[0].Name, commandline.LogLevel.Name, zapcore.WarnLevel.String())

	err := app.Run([]string{os.Args[0], app.Commands[0].Name, commandline.LogLevel.Name, zapcore.WarnLevel.String()})
	if err != nil {
		panic(err)
	}

	// output: running ./example test --log-level warn
}
