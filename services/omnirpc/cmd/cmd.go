package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/commandline"
	rpcConfig "github.com/synapsecns/sanguine/serivces/omnirpc/config"
	"github.com/synapsecns/sanguine/serivces/omnirpc/latency"
	"github.com/synapsecns/sanguine/serivces/omnirpc/proxy"
	"github.com/synapsecns/sanguine/serivces/omnirpc/rpcmap"
	"github.com/synapsecns/synapse-node/config"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const appName = "omnirpc"

// Start starts the command line.
func Start(args []string) {
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
				rpcMap, err := rpcmap.GetPublicRPCMap(c.Context)
				if err != nil {
					return fmt.Errorf("could not get rpc map: %w", err)
				}

				res := latency.GetRPCLatency(c.Context, time.Second*5, rpcMap.RawMap()[c.Int("chain-id")])
				DisplayLatency(res)

				return nil
			},
		},
		{
			Name:  "chainlist-server",
			Usage: "runs a chainlist proxy server",
			Flags: []cli.Flag{portFlag},
			Action: func(c *cli.Context) error {
				rpcMap, err := rpcmap.GetPublicRPCMap(c.Context)
				if err != nil {
					return fmt.Errorf("could not get rpc map: %w", err)
				}

				server := proxy.NewProxy(uint32(c.Int("port")), rpcMap)

				server.Run(c.Context)

				return nil
			},
		},
		{
			Name:  "public-config",
			Usage: "output a public config file from chainlist.org",
			Flags: []cli.Flag{outputFlag},
			Action: func(c *cli.Context) error {
				rpcMap, err := rpcmap.GetPublicRPCMap(c.Context)
				if err != nil {
					return fmt.Errorf("could not get rpc map: %w", err)
				}

				output := rpcConfig.MarshallFromMap(rpcMap)
				outputConfig, err := os.Create(c.String(outputFlag.Name))
				if err != nil {
					return fmt.Errorf("could not create config file: %w", err)
				}

				defer func() {
					_ = outputConfig.Close()
				}()

				_, err = outputConfig.WriteString(output)
				if err != nil {
					return fmt.Errorf("could not write to file: %w", err)
				}

				fmt.Printf("written to %s \n", c.String(outputFlag.Name))
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "run a server from a config",
			Flags: []cli.Flag{
				configFlag,
				portFlag,
			},
			Action: func(c *cli.Context) error {
				rpcMap, err := rpcConfig.UnmarshallConfigFromFile(c.String(configFlag.Name))
				if err != nil {
					return fmt.Errorf("could not unmarshall config: %w", err)
				}

				server := proxy.NewProxy(uint32(c.Int("port")), rpcMap)

				server.Run(c.Context)

				return nil
			},
		},
	}
	shellCommand := commandline.GenerateShellCommand(app.Commands)
	app.Commands = append(app.Commands, shellCommand)
	app.Action = shellCommand.Action

	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
