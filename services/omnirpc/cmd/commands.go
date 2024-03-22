package cmd

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/omnirpc/modules/confirmedtofinalized"
	"github.com/synapsecns/sanguine/services/omnirpc/modules/harmonyproxy"
	"os"
	"time"

	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core"
	rpcConfig "github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/debug"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"github.com/urfave/cli/v2"
)

var latencyCommand = &cli.Command{
	Name:  "check-latency",
	Usage: "checks latency for all rpc endpoints known for a chain id",
	Flags: []cli.Flag{chainIDFlag},
	Action: func(c *cli.Context) error {
		rConfig, err := rpcConfig.GetPublicRPCConfig()
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		chainConfig, ok := rConfig.Chains[uint32(c.Int(chainIDFlag.Name))]
		if !ok {
			return fmt.Errorf("could not get chain config for chain %d", c.Int(chainIDFlag.Name))
		}

		res := rpcinfo.GetRPCLatency(c.Context, time.Second*5, chainConfig.RPCs, metrics.Get())
		DisplayLatency(res)

		return nil
	},
}

var chainListCommand = &cli.Command{
	Name:  "chainlist-server",
	Usage: "runs a chainlist proxy server",
	Flags: []cli.Flag{portFlag},
	Action: func(c *cli.Context) error {
		// Create a large heap allocation of 10 GiB
		// See: https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap/
		_ = make([]byte, 10<<30)

		rConfig, err := rpcConfig.GetPublicRPCConfig()
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		if c.Int(portFlag.Name) != 0 {
			rConfig.Port = uint16(c.Int(portFlag.Name))
		}

		if rConfig.Port == 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		server := proxy.NewProxy(rConfig, metrics.Get())

		server.Run(c.Context)

		return nil
	},
}

var publicConfigCommand = &cli.Command{
	Name:  "public-config",
	Usage: "output a public config file from chainlist.org",
	Flags: []cli.Flag{outputFlag},
	Action: func(c *cli.Context) error {
		rConfig, err := rpcConfig.GetPublicRPCConfig()
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		output, err := rConfig.Marshall()
		if err != nil {
			return fmt.Errorf("could not get rpc map: %w", err)
		}

		outputConfig, err := os.Create(core.ExpandOrReturnPath(c.String(outputFlag.Name)))
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
		// Create a large heap allocation of 10 GiB
		// See: https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap/
		_ = make([]byte, 10<<30)

		fileContents, err := os.ReadFile(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not read file %s: %w", c.String(configFlag.Name), err)
		}
		rConfig, err := rpcConfig.UnmarshallConfig(fileContents)
		if err != nil {
			return fmt.Errorf("could not unmarshall config: %w", err)
		}

		if c.Int(portFlag.Name) != 0 {
			rConfig.Port = uint16(c.Int(portFlag.Name))
		}

		if rConfig.Port == 0 {
			rConfig.Port = uint16(freeport.GetPort())
		}

		server := proxy.NewProxy(rConfig, metrics.Get())

		server.Run(c.Context)

		return nil
	},
}

var debugResponse = &cli.Command{
	Name:  "debug-response",
	Usage: "used for debugging responses and finding diff between rpcs",
	Flags: []cli.Flag{
		fileFlag,
	},
	Action: func(c *cli.Context) error {
		diffFile, err := os.ReadFile(c.String(fileFlag.Name))
		if err != nil {
			return fmt.Errorf("could not read file: %w", err)
		}
		//nolint:wrapcheck
		return debug.HashDiff(diffFile)
	},
}

// latestRewrite rewrites latest block numbers for a single rpc url.
var latestRewrite = &cli.Command{
	Name:  "latest-rewrite",
	Usage: "A simple rpc proxy for one-off integration tests. Rewrites block queries that use \"latest\" to \"finalized\"",
	Flags: []cli.Flag{
		rpcFlag,
		portFlag,
	},
	Action: func(c *cli.Context) error {
		simpleProxy := confirmedtofinalized.NewProxy(c.String(rpcFlag.Name), metrics.Get(), c.Int(portFlag.Name))

		err := simpleProxy.Run(c.Context)
		if err != nil {
			return fmt.Errorf("return err: %w", err)
		}
		return nil
	},
}

var harmonyProxy = &cli.Command{
	Name:  "harmony-confirm",
	Usage: "An experimental harmony confirmation client",
	Flags: []cli.Flag{
		rpcFlag,
		portFlag,
	},
	Action: func(c *cli.Context) error {
		simpleProxy := harmonyproxy.NewHarmonyProxy(c.String(rpcFlag.Name), metrics.Get(), c.Int(portFlag.Name))

		err := simpleProxy.Run(c.Context)
		if err != nil {
			return fmt.Errorf("return err: %w", err)
		}
		return nil
	},
}
