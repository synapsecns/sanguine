package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/api"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer/node"
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
var clickhouseAddressFlag = &cli.StringFlag{
	Name:     "address",
	Usage:    "--address <path/to/database> or <database url>",
	Value:    "",
	Required: true,
}
var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
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
var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	//Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}

		db, err := api.InitDB(c.Context, c.String(clickhouseAddressFlag.Name))
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}
		clients := make(map[uint32]bind.ContractBackend)
		for _, client := range decodeConfig.Chains {
			backendClient, err := ethclient.DialContext(c.Context, client.RPCURL)
			if err != nil {
				return fmt.Errorf("could not start client for %s", client.RPCURL)
			}
			clients[client.ChainID] = backendClient
		}
		explorerBackfiller, err := node.NewExplorerBackfiller(db, decodeConfig, clients)
		if err != nil {
			return fmt.Errorf("could not create explorer backfiller: %w", err)
		}
		err = explorerBackfiller.Backfill(c.Context)
		if err != nil {
			return fmt.Errorf("could not backfill backfiller: %w", err)
		}
		return nil
	},
}

func init() {
	portFlag.Value = uint(freeport.GetPort())
}

//
//// get the ABI for each contract
// bridgeConfigABI, err := bridgeconfig.BridgeConfigV3MetaData.GetAbi()
// if err != nil || bridgeConfigABI == nil {
// return nil, fmt.Errorf("could not get bridge config v3 abi: %w", err)
//}
// swapABI, err := swap.SwapFlashLoanMetaData.GetAbi()
// if err != nil || swapABI == nil {
// return nil, fmt.Errorf("could not get swap flash loan abi: %w", err)
//}
//bridgeABI, err := bridge.SynapseBridgeMetaData.GetAbi()
//if err != nil || bridgeABI == nil {
//return nil, fmt.Errorf("could not get bridge abi: %w", err)
//}
//
//// create the client
//ethClient, err := ethclient.DialContext(ctx, chainConfig.RPCURL)
//// create the consumer Fetcher
//fetcher := consumer.NewFetcher(gqlClient.NewClient(http.DefaultClient, baseURL))
//// create the bridge config Fetcher
//bridgeConfigFetcher, err := consumer.NewBridgeConfigFetcher(common.HexToAddress(chainConfig.BridgeConfigV3Address), ethClient)
//if err != nil || bridgeConfigFetcher == nil {
//return nil, fmt.Errorf("could not create bridge config Fetcher: %w", err)
//}
//// create the bridge parser
//bridgeParser, err := consumer.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.SynapseBridgeAddress), *bridgeConfigFetcher, fetcher)
//if err != nil || bridgeParser == nil {
//return nil, fmt.Errorf("could not create bridge parser: %w", err)
//}
//// create the swap parsers
//swapParsers := make(map[common.Address]*consumer.SwapParser)
//for _, swapAddress := range chainConfig.SwapFlashLoanAddresses {
//// create the swap Fetcher
//swapFetcher, err := consumer.NewSwapFetcher(common.HexToAddress(swapAddress), ethClient)
//if err != nil || swapFetcher == nil {
//return nil, fmt.Errorf("could not create swap Fetcher: %w", err)
//}
//swapParser, err := consumer.NewSwapParser(consumerDB, common.HexToAddress(swapAddress), *swapFetcher, fetcher)
//if err != nil || swapParser == nil {
//return nil, fmt.Errorf("could not create swap parser: %w", err)
//}
//swapParsers[common.HexToAddress(swapAddress)] = swapParser
//}
//chainBackfiller := backfill.NewChainBackfiller(consumerDB, bridgeParser, swapParsers, *fetcher, chainConfig)
//
//return chainBackfiller, nil
