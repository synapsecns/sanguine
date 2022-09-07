package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"


	"github.com/hashicorp/consul/sdk/freeport"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
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


//placeholder
var placeholderCommand = &cli.Command{
	Name:        "gm",
	Description: "gm",
	Usage:       "gm",
	Action: func(c *cli.Context) error {
		fmt.Println("gm")
		return nil
	},
}

func init() {
	ports := freeport.Get(1)
	if len(ports) > 0 {
		// portFlag.Value = uint(ports[0])
	}
}
