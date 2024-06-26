// Package cmd provides the command line interface for the opbot.
package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	exampleConfig "github.com/synapsecns/sanguine/ethergo/examples/signer-example/config"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
)

// Start starts the command line tool.
func Start(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()

	app.Description = buildInfo.VersionString() + "Opbot is a tool to manage operations."
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.EnableBashCompletion = true
	app.Before = func(c *cli.Context) error {
		// nolint:wrapcheck
		return metrics.Setup(c.Context, buildInfo)
	}
	app.Flags = []cli.Flag{fileFlag}
	app.Action = action
	err := app.Run(args)

	if err != nil {
		panic(err)
	}
}

var fileFlag = &cli.StringFlag{
	Name:     "config",
	Usage:    "--config /Users/synapsecns/config.yaml",
	Required: true,
}

var action = func(c *cli.Context) error {
	configFile, err := os.ReadFile(c.String(fileFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	var cfg exampleConfig.ExampleConfig
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if ok, err := cfg.SignerConfig.IsValid(c.Context); !ok {
		panic(fmt.Errorf("failed to validate signer config: %w", err))
	}

	signer, err := signerConfig.SignerFromConfig(c.Context, cfg.SignerConfig)
	if err != nil {
		return fmt.Errorf("could not create signer: %w", err)
	}

	fmt.Printf("signer address is %s", signer.Address())

	return nil
}
