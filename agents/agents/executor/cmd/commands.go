package cmd

import (
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// inforCommand gets info about using the executor service.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use executor cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}
