package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/synapsecns/sanguine/services/scribe/server"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql"
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
	Usage:     "--config /Users/synapsecns/config.toml",
	TakesFile: true,
}

var backfillCommand = &cli.Command{
	Name:        "backfill",
	Description: "backfills up to a block and then halts",
	Usage:       "backfill --config /path/to/config.toml",
	Flags:       []cli.Flag{configFlag},
	Action: func(c *cli.Context) error {
		decodeConfig, err := config.DecodeConfig(c.String(configFlag.Name))
		if err != nil {
			return fmt.Errorf("could not decode config: %w", err)

		}

		// TODO: this should be done in a node folder
		// temporary for now, TODO add a full config
		tempDir, err := os.MkdirTemp("", "")
		if err != nil {
			return fmt.Errorf("could not create temp dir: %w", err)
		}

		fmt.Println("temp:")
		fmt.Println(tempDir)

		db, err := sql.NewStoreFromConfig(c.Context, dbcommon.Sqlite, tempDir)
		if err != nil {
			return fmt.Errorf("could not create store: %w", err)
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

		scribeBackfiller, err := backfill.NewScribeBackfiller(db, clients, *decodeConfig)
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

var portFlag = &cli.UintFlag{
	Name:  "port",
	Usage: "--port 5121",
	Value: uint(freeport.Get(1)[0]),
}

var serverCommand = &cli.Command{
	Name:        "server",
	Description: "starts a graphql server",
	Flags:       []cli.Flag{portFlag},
	Action: func(c *cli.Context) error {
		return server.Start(uint16(c.Uint(portFlag.Name)))
	},
}
