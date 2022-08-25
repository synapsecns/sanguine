// Package main contains a generator for copying files exported files from geth
// in order to use private fields. The resulting files should not be modified directly
// but if there are new methods you need exported, generators, etc that can be done in other files
// that will now have access to the private fields. These generated files should only be used for testing
//
// TODO: look into implementing a tag for tests in order to make sure nothing in testutils/ is used in a production build
// we haven't done this yet because of the poor ux in an ide as far as having to add a `-tag`.
package main

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/services/omnirpc/pkg"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const appName = "modulecopier"

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = "0.1.0"
	app.Description = "Used for checking the lowest latency rpc endpoint fora given chain"
	app.Commands = []*cli.Command{
		{
			Name:  "check-latency",
			Usage: "checks latency for all rpc endpoints known for a chain id",
			Flags: []cli.Flag{chainIDFlag},
			Action: func(c *cli.Context) error {
				rpcMap, err := pkg.GetRPCMap(c.Context)
				if err != nil {
					return fmt.Errorf("could not get rpc map: %w", err)
				}

				res := pkg.GetRPCLatency(c.Context, time.Second*5, rpcMap[c.Int("chain-id")])
				DisplayLatency(res)

				return nil
			},
		},
	}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
