package cmd

import (
	"fmt"
	"github.com/phayes/freeport"
	rpcConfig "github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var latencyCommand = &cli.Command{
	Name:  "check-latency",
	Usage: "checks latency for all rpc endpoints known for a chain id",
	Flags: []cli.Flag{chainIDFlag},
	Action: func(c *cli.Context) error {
		rConfig, err := rpcConfig.GetPublicRPCConfig(c.Context)
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		chainConfig, ok := rConfig.Chains[uint32(c.Int(chainIDFlag.Name))]
		if !ok {
			return fmt.Errorf("could not get chain config for chain %d", c.Int(chainIDFlag.Name))
		}

		res := rpcinfo.GetRPCLatency(c.Context, time.Second*5, chainConfig.RPCs)
		DisplayLatency(res)

		return nil
	},
}

var chainListCommand = &cli.Command{
	Name:  "chainlist-server",
	Usage: "runs a chainlist proxy server",
	Flags: []cli.Flag{portFlag},
	Action: func(c *cli.Context) error {
		rConfig, err := rpcConfig.GetPublicRPCConfig(c.Context)
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		if c.Int(portFlag.Name) != 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		if rConfig.Port == 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		server := proxy.NewProxy(rConfig)

		server.Run(c.Context)

		return nil
	},
}

var publicConfigCommand = &cli.Command{
	Name:  "public-config",
	Usage: "output a public config file from chainlist.org",
	Flags: []cli.Flag{outputFlag},
	Action: func(c *cli.Context) error {
		rConfig, err := rpcConfig.GetPublicRPCConfig(c.Context)
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		output, err := rConfig.Marshall()
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		outputConfig, err := os.Create(c.String(outputFlag.Name))
		if err != nil {
			return fmt.Errorf("could not create config file: %w", err)
		}

		defer func() {
			_ = outputConfig.Close()
		}()

		_, err = outputConfig.Write(output)
		if err != nil {
			return fmt.Errorf("could not write to file: %w", err)
		}

		fmt.Printf("written to %s \n", c.String(outputFlag.Name))
		return nil
	},
}

var serverCommand = &cli.Command{
	Name:  "server",
	Usage: "run a server from a config",
	Flags: []cli.Flag{
		configFlag,
		portFlag,
	},
	Action: func(c *cli.Context) error {
		fileContents, err := os.ReadFile(c.String(configFlag.Name))
		if err != nil {
			return fmt.Errorf("could not read file %s: %w", c.String(configFlag.Name), err)
		}
		rConfig, err := rpcConfig.UnmarshallConfig(fileContents)
		if err != nil {
			return fmt.Errorf("could not unmarshall config: %w", err)
		}

		if c.Int(portFlag.Name) != 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		if rConfig.Port == 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		server := proxy.NewProxy(rConfig)

		server.Run(c.Context)

		return nil
	},
}
