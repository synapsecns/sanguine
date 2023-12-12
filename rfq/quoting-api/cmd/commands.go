package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest"
	"github.com/urfave/cli/v2"
)

// infoCommand gets info the quoter API.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "quoter help",
	Action: func(c *cli.Context) error {
		fmt.Println("run quoter --config /path/to/config.yaml to start the quoter")
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Path/To/Config.yaml",
	TakesFile: true,
	Required:  true,
}

var quoterCommand = &cli.Command{
	Name:        "quoter",
	Description: "runs the quoter server",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		cfg, err := config.LoadConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			fmt.Println("Error loading config:", err)
		}
		fmt.Printf("Config loaded: %+v\n", cfg)
		restAPI, err := rest.NewRestAPIServer(c.Context, &cfg)
		if err != nil {
			return fmt.Errorf("could not create rest api server: %w", err)
		}
		restAPI.Setup()
		err = restAPI.Run()
		if err != nil {
			return fmt.Errorf("could not run rest api server: %w", err)
		}
		return nil
	},
}
