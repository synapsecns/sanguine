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
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/explorer/api"
	indexerconfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
	serverconfig "github.com/synapsecns/sanguine/services/explorer/config/server"
	"github.com/synapsecns/sanguine/services/explorer/node"
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

var clickhouseAddressFlag = &cli.StringFlag{
	Name:     "address",
	Usage:    "--address pass 'default' to use the default clickhouse address",
	Value:    "",
	Required: false,
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
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		fmt.Println("port", c.Uint("port"))
		decodeConfig, err := serverconfig.DecodeServerConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)
		}

		err = api.Start(c.Context, decodeConfig, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}

		return nil
	},
}

// nolint:dupl
var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	Flags:       []cli.Flag{configFlag, clickhouseAddressFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := indexerconfig.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}
		db, err := api.InitDB(c.Context, c.String(clickhouseAddressFlag.Name), false, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}
		clients := make(map[uint32]bind.ContractBackend)
		for _, client := range decodeConfig.Chains {
			backendClient, err := ethclient.DialContext(c.Context, decodeConfig.RPCURL+fmt.Sprintf("%d", client.ChainID))
			if err != nil {
				return fmt.Errorf("could not start client for %s", client.RPCURL)
			}
			clients[client.ChainID] = backendClient
		}
		explorerBackfiller, err := node.NewExplorerBackfiller(db, decodeConfig, clients, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not create explorer backfiller: %w", err)
		}
		err = explorerBackfiller.Backfill(c.Context, false)
		if err != nil {
			return fmt.Errorf("could not backfill backfiller: %w", err)
		}
		return nil
	},
}

// nolint:dupl
var livefillCommand = &cli.Command{
	Name:        "livefill",
	Description: "livefills explorer",
	Flags:       []cli.Flag{configFlag, clickhouseAddressFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := indexerconfig.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}
		db, err := api.InitDB(c.Context, c.String(clickhouseAddressFlag.Name), false, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}
		clients := make(map[uint32]bind.ContractBackend)
		for _, client := range decodeConfig.Chains {
			backendClient, err := ethclient.DialContext(c.Context, decodeConfig.RPCURL+fmt.Sprintf("%d", client.ChainID))
			if err != nil {
				return fmt.Errorf("could not start client for %s", client.RPCURL)
			}
			clients[client.ChainID] = backendClient
		}
		explorerBackfiller, err := node.NewExplorerBackfiller(db, decodeConfig, clients, metrics.Get())
		if err != nil {
			return fmt.Errorf("could not create explorer backfiller: %w", err)
		}
		err = explorerBackfiller.Backfill(c.Context, true)
		if err != nil {
			return fmt.Errorf("could not backfill backfiller: %w", err)
		}
		return nil
	},
}

func init() {
	portFlag.Value = uint(freeport.GetPort())
}
