package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/urfave/cli/v2"
)

func Run(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()
	app.Version = buildInfo.Version()
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.Commands = []*cli.Command{
		{
			Name:        "ignore-generated",
			Description: "removes autogenered files from coverage report",
			Action: func(c *cli.Context) error {
				// TODO: implement
				panic("not implemented")
			},
		},
	}
}
