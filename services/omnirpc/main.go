package main

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	"github.com/synapsecns/sanguine/serivces/omnirpc/latency"
	"github.com/synapsecns/sanguine/serivces/omnirpc/rpcmap"
	"github.com/synapsecns/synapse-node/config"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const appName = "omnirpc"

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = config.AppVersion
	app.Description = "Used for checking the lowest latency rpc endpoint fora given chain"
	app.Commands = []*cli.Command{
		{
			Name:  "check-latency",
			Usage: "checks latency for all rpc endpoints known for a chain id",
			Flags: []cli.Flag{chainIDFlag},
			Action: func(c *cli.Context) error {
				rpcMap, err := rpcmap.GetRPCMap(c.Context)
				if err != nil {
					return fmt.Errorf("could not get rpc map: %w", err)
				}

				res := latency.GetRPCLatency(c.Context, time.Second*5, rpcMap[c.Int("chain-id")])
				DisplayLatency(res)

				return nil
			},
		}}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
