package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// infoComand gets info about using the scribe service.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use scribe cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

var portFlag = &cli.UintFlag{
	Name:  "port",
	Usage: "--port 5121",
	Value: 0,
}

var grpcPortFlag = &cli.UintFlag{
	Name:  "grpc-port",
	Usage: "--port 5121",
	Value: 0,
}

var dbFlag = &cli.StringFlag{
	Name:     "db",
	Usage:    "--db <sqlite> or <mysql>",
	Value:    "sqlite",
	Required: true,
}

var pathFlag = &cli.StringFlag{
	Name:     "path",
	Usage:    "--path <path/to/database> or <database url>",
	Value:    "",
	Required: true,
}

var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	Flags:       []cli.Flag{configFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}

		db, err := api.InitDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name))
		if err != nil {
			return fmt.Errorf("could not initialize database: %w", err)
		}

		clients := make(map[uint32]backfill.ScribeBackend)
		// TODO: should be resistant to errors on startup from a single chain
		for _, client := range decodeConfig.Chains {
			backendClient, err := ethclient.DialContext(c.Context, client.RPCUrl)
			if err != nil {
				return fmt.Errorf("could not start client for %s", client.RPCUrl)
			}

			clients[client.ChainID] = backendClient
		}

		scribeBackfiller, err := backfill.NewScribeBackfiller(db, clients, decodeConfig)
		if err != nil {
			return fmt.Errorf("could not create scribe backfiller: %w", err)
		}

		err = scribeBackfiller.Backfill(c.Context)
		if err != nil {
			return fmt.Errorf("could not backfill backfiller: %w", err)
		}

		return nil
	},
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{portFlag, dbFlag, pathFlag},
	Action: func(c *cli.Context) error {
		err := api.Start(c.Context, api.Config{
			HTTPPort: uint16(c.Uint(portFlag.Name)),
			Database: c.String(dbFlag.Name),
			Path:     c.String(pathFlag.Name),
			GRPCPort: uint16(c.Uint(grpcPortFlag.Name)),
		})
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}

		return nil
	},
}

func init() {
	ports := freeport.Get(1)
	if len(ports) > 0 {
		portFlag.Value = uint(ports[0])
	}
}
