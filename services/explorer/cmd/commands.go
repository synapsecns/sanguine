package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/services/explorer/api"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// infoCommand references the help info from the cmd.md file and presents it.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use explorer cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var portFlag = &cli.UintFlag{
	Name:  "port",
	Usage: "--port 5121",
	Value: 0,
}

var addressFlag = &cli.StringFlag{
	Name:     "address",
	Usage:    "--address <address>",
	Value:    "",
	Required: true,
}

var scribeURL = &cli.StringFlag{
	Name:     "scribe-url",
	Usage:    "--scribe-url <scribe-url>",
	Required: true,
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{portFlag, addressFlag},
	Action: func(c *cli.Context) error {
		fmt.Println("port", c.Uint("port"))
		err := api.Start(c.Context, api.Config{
			HTTPPort:  uint16(c.Uint(portFlag.Name)),
			Address:   c.String(addressFlag.Name),
			ScribeURL: c.String(scribeURL.Name),
		})
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}

		return nil
	},
}

func init() {
	portFlag.Value = uint(freeport.GetPort())
}
