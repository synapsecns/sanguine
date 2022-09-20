package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"

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
